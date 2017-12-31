package bog

import "fmt"

type Contact struct {
	ID        uint
	ClientID  uint
	FirstName string
	LastName  string
	Email     string
	IsPrimary bool
	service   ContactService
}

// An interface for Contact model DAO operations
type ContactService interface {
	// Get a single Contact
	Get(id uint) (Contact, error)

	// Get all contacts for a client
	GetAll(clientId uint) ([]Contact, error)

	// Get the primary contact for a client
	GetPrimary(clientId uint) (Contact, error)

	// Create a new Contact.
	Insert(contact *Contact) error

	// Update an existing Contact in the database.
	Update(contact *Contact) error

	// Delete an existing Contact in the database.
	Delete(id uint) error
}

type contactServiceFactory func() ContactService

var NewContactService contactServiceFactory

// Create an empty Contact instance
func NewContact() Contact {
	service := NewContactService()
	c := Contact{
		service: service,
	}

	return c
}

// Get a Contact instance matching id
func LoadContact(id uint) (Contact, error) {
	service := NewContactService()

	if id == 0 {
		return Contact{}, fmt.Errorf("Cannot load contact: id = 0")
	}

	return service.Get(id)
}

// Contact.Insert inserts the contact.
func (c *Contact) Insert() error {
	if c.ID == 0 {
		return c.service.Insert(c)
	}

	return fmt.Errorf("Cannot inset contact if ID > 0")
}

// Contact.Save inserts or updates the Contact.
func (c *Contact) Update() error {
	if c.ID == 0 {
		return fmt.Errorf("Cannot update Contact if ID = 0")
	}

	return c.service.Update(c)
}

// Contact.Delete deletes the client.
func (c *Contact) Delete() error {
	if c.ID == 0 {
		return fmt.Errorf("cannot delete Contact id = 0")
	}

	return c.service.Delete(c.ID)
}
