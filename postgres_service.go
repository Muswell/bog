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
	DB *sql.DB
}

func init() {
	log.Println("Running postgres-service init")
	NewClientService = func() ClientService {
		s := PostgresClientService{}
		db, err := database.New()

		if err != nil {
			log.Fatalf("Could not initialixe PostgresClientService: %v", err)
		}
		s.DB = db

		return &s
	}
}

// Get a client from the clients map
func (s *PostgresClientService) Get(id uint) (Client, error) {
	c := Client{}

	if id == 0 {
		return c, fmt.Errorf("PostgresClientService.Get requires a valid id.")
	}

	db := s.DB

	defer db.Close()

	row := db.QueryRow("SELECT id, name, address_1, address_2, city, state, zip FROM clients WHERE id = $1", id)

	var name string
	var address1 string
	var address2 string
	var city string
	var state string
	var zip string

	err := row.Scan(&id, &name, &address1, &address2, &city, &state, &zip)
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

// Get all clients from the clients map
func (s *PostgresClientService) GetAll() ([]Client, error) {
	clients := []Client{}

	rows, err := s.DB.Query("SELECT id, name, address_1, address_2, city, state, zip FROM clients")

	if err != nil {
		return make([]Client, 0), err
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

		clients = append(clients, c)
	}

	return clients, nil
}

// Add a client to the clients map
func (s *PostgresClientService) Insert(c *Client) error {
	if c.ID != 0 {
		return fmt.Errorf("Can't insert client: id is not zero")
	}

	var id uint
	err := s.DB.QueryRow(
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

// Update a client in the clients map
func (s *PostgresClientService) Update(c *Client) error {
	if c.ID == 0 {
		return fmt.Errorf("Cannot update a client which has not been inserted.")
	}

	_, err := s.DB.Exec(
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

// Remove a client from the clients map
func (s *PostgresClientService) Delete(id uint) error {
	_, err := s.DB.Exec("DELETE FROM clients WHERE id = $1", id)

	if err == nil {
		log.Printf("Client deleted name: id: %d", id)
	}

	return err
}
