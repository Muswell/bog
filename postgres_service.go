// +build postgres_service

package bog

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/muswell/bog/database"
	"log"
)

// PostgreSQL implementation of ClientService
type PostgresClientService struct {
	DB func() (*sql.DB, error)
}

// PostgreSQL implementation of ContactService
type PostgresContactService struct {
	DB func() (*sql.DB, error)
}

func init() {
	log.Println("Running postgres-service init")
	NewClientService = func() ClientService {
		s := PostgresClientService{
			DB: database.New,
		}

		return &s
	}

	NewContactService = func() ContactService {
		s := PostgresContactService{
			DB: database.New,
		}

		return &s
	}
}

/*---------------------------------------------------
 * Client Service
 *-------------------------------------------------*/
// Get a client from the clients table
func (s *PostgresClientService) Get(id uint) (Client, error) {
	c := Client{
		service: s,
	}

	if id == 0 {
		return c, fmt.Errorf("PostgresClientService.Get requires a valid id.")
	}

	db, err := s.DB()
	if err != nil {
		return c, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name, address_1, address_2, city, state, zip FROM clients WHERE id = $1", id)

	var name string
	var address1 string
	var address2 string
	var city string
	var state string
	var zip string

	err = row.Scan(&id, &name, &address1, &address2, &city, &state, &zip)
	if err != nil {
		return c, err
	}

	c.ID = id
	c.Name = name
	c.Address1 = address1
	c.Address2 = address2
	c.City = city
	c.State = state
	c.Zip = zip

	return c, nil
}

// Get all clients from the clients table
func (s *PostgresClientService) GetAll() ([]Client, error) {
	clients := []Client{}

	db, err := s.DB()
	if err != nil {
		return clients, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, address_1, address_2, city, state, zip FROM clients")

	if err != nil {
		return clients, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var name string
		var address1 string
		var address2 string
		var city string
		var state string
		var zip string

		err = rows.Scan(&id, &name, &address1, &address2, &city, &state, &zip)
		if err != nil {
			return clients, err
		}

		c := Client{}
		c.ID = id
		c.Name = name
		c.Address1 = address1
		c.Address2 = address2
		c.City = city
		c.State = state
		c.Zip = zip
		c.service = s

		clients = append(clients, c)
	}

	return clients, nil
}

// Add a client to the clients table
func (s *PostgresClientService) Insert(c *Client) error {
	if c.ID != 0 {
		return fmt.Errorf("Can't insert client: id is not zero")
	}

	var id uint

	db, err := s.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(
		"INSERT INTO clients (name, address_1, address_2, city, state, zip) "+
			"VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		c.Name,
		c.Address1,
		c.Address2,
		c.City,
		c.State,
		c.Zip,
	).Scan(&id)
	if err != nil {
		return err
	}

	c.ID = id
	log.Printf("Client inserted name: %s, id: %d", c.Name, c.ID)
	return nil
}

// Update a client in the clients table
func (s *PostgresClientService) Update(c *Client) error {
	if c.ID == 0 {
		return fmt.Errorf("Cannot update a client which has not been inserted.")
	}

	db, err := s.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"UPDATE clients set name = $1, address_1 = $2, address_2 = $3, city = $4, state = $5, zip = $6 WHERE id = $7",
		c.Name,
		c.Address1,
		c.Address2,
		c.City,
		c.State,
		c.Zip,
		c.ID,
	)

	if err == nil {
		log.Printf("Client updated name: %s, id: %d", c.Name, c.ID)
	}

	return err
}

// Remove a client from the clients table
func (s *PostgresClientService) Delete(id uint) error {
	db, err := s.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM clients WHERE id = $1", id)

	if err == nil {
		log.Printf("Client deleted name: id: %d", id)
	}

	return err
}

/*---------------------------------------------------
 * Contact Service
 *-------------------------------------------------*/
// Get a Contact from the contacts table
func (s *PostgresContactService) Get(id uint) (Contact, error) {
	c := Contact{
		service: s,
	}

	if id == 0 {
		return c, fmt.Errorf("PostgresContactService.Get requires a valid id.")
	}

	db, err := s.DB()
	if err != nil {
		return c, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, client_id, first_name, last_name, email, is_primary FROM contacts WHERE id = $1", id)

	var clientId uint
	var firstName string
	var lastName string
	var email string
	var isPrimary bool

	err = row.Scan(&id, &clientId, &firstName, &lastName, &email, &isPrimary)
	if err != nil {
		return c, err
	}

	c.ID = id
	c.ClientID = clientId
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	c.IsPrimary = isPrimary

	return c, nil
}

// Get all contact for a client
func (s *PostgresContactService) GetAll(clientId uint) ([]Contact, error) {
	contacts := []Contact{}

	if clientId == 0 {
		return contacts, fmt.Errorf("Can not load contacts, clientId = 0")
	}

	db, err := s.DB()
	if err != nil {
		return contacts, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, client_id, first_name, last_name, email, is_primary FROM contacts WHERE client_id = $1", clientId)

	if err != nil {
		return contacts, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var clientId uint
		var firstName string
		var lastName string
		var email string
		var isPrimary bool

		err := rows.Scan(&id, &clientId, &firstName, &lastName, &email, &isPrimary)
		if err != nil {
			return contacts, err
		}

		c := Contact{}
		c.ID = id
		c.ClientID = clientId
		c.FirstName = firstName
		c.LastName = lastName
		c.Email = email
		c.IsPrimary = isPrimary
		c.service = s

		contacts = append(contacts, c)
	}

	return contacts, nil
}

// Get primary contact for a client
func (s *PostgresContactService) GetPrimary(clientId uint) (Contact, error) {
	c := Contact{
		service: s,
	}

	if clientId == 0 {
		return c, fmt.Errorf("PostgresContactService.GetPrimary requires a valid clientId.")
	}

	db, err := s.DB()
	if err != nil {
		return c, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, client_id, first_name, last_name, email, is_primary FROM contacts WHERE client_id = $1 AND is_primary = $2 LIMIT 1", clientId, true)

	var id uint
	var firstName string
	var lastName string
	var email string
	var isPrimary bool

	err = row.Scan(&id, &clientId, &firstName, &lastName, &email, &isPrimary)
	if err != nil {
		return c, err
	}

	c.ID = id
	c.ClientID = clientId
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	c.IsPrimary = isPrimary

	return c, nil
}

// Add a contact to the contacts table
func (s *PostgresContactService) Insert(c *Contact) error {
	if c.ID != 0 {
		return fmt.Errorf("Can't insert contact: id is not zero")
	}

	db, err := s.DB()
	if err != nil {
		return err
	}

	defer db.Close()

	if c.IsPrimary {
		err := clearExistingPrimaryContact(c.ClientID)
		if err != nil {
			return err
		}
	}

	var id uint
	err = db.QueryRow(
		"INSERT INTO contacts (client_id, first_name, last_name, email, is_primary) "+
			"VALUES($1, $2, $3, $4, $5) RETURNING id",
		c.ClientID,
		c.FirstName,
		c.LastName,
		c.Email,
		c.IsPrimary,
	).Scan(&id)
	if err != nil {
		return err
	}

	c.ID = id
	log.Printf("Contact inserted %v", c)
	return nil
}

// Set all contacts is_primary to false for a given clientId
func clearExistingPrimaryContact(clientId uint) error {
	if clientId == 0 {
		return fmt.Errorf("clearExistingPrimaryContact requires a valid clientId.")
	}

	db, err := database.New()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(
		"UPDATE contacts set is_primary = $1 WHERE client_id = $2",
		false,
		clientId,
	)

	return err
}

// Update a contact in the contacts table
func (s *PostgresContactService) Update(c *Contact) error {
	if c.ID == 0 {
		return fmt.Errorf("Cannot update a contact which has not been inserted.")
	}

	db, err := s.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	if c.IsPrimary {
		err := clearExistingPrimaryContact(c.ClientID)
		if err != nil {
			return err
		}
	}

	// cannot update client_id
	_, err = db.Exec(
		"UPDATE contacts set first_name = $1, last_name = $2, email = $3, is_primary = $4 WHERE id = $5",
		c.FirstName,
		c.LastName,
		c.Email,
		c.IsPrimary,
		c.ID,
	)

	if err == nil {
		log.Printf("contact updated id: %d", c.ID)
	}

	return err
}

// Remove a contact from the contacts table
func (s *PostgresContactService) Delete(id uint) error {
	db, err := s.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM contacts WHERE id = $1", id)

	if err == nil {
		log.Printf("contact deleted name: id: %d", id)
	}

	return err
}
