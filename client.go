package bog

import (
	"fmt"
)

// Model for the client
type Client struct {
	ID       uint
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
	service  ClientService
}

// An interface for Client model DAO operations
type ClientService interface {
	// Get a single Client
	Get(id uint) (Client, error)

	// Get all clients
	GetAll() ([]Client, error)

	// Create a new Client.
	Insert(client *Client) error

	// Update an existing Client in the database.
	Update(client *Client) error

	// Delete an existing Client in the database.
	Delete(id uint) error
}

type clientServiceFactory func() ClientService

var NewClientService clientServiceFactory

func NewClient() Client {
	service := NewClientService()
	c := Client{
		service: service,
	}

	return c
}

func GetClient(id uint) (Client, error) {
	service := NewClientService()

	if id == 0 {
		return Client{}, fmt.Errorf("Cannot load client: id = 0")
	}

	return service.Get(id)
}

func GetAllClients() ([]Client, error) {
	service := NewClientService()

	return service.GetAll()
}

// Client.Insert inserts the client.
func (c *Client) Insert() error {
	if c.ID == 0 {
		return c.service.Insert(c)
	}

	return fmt.Errorf("Cannot inset client if ID > 0")
}

// Client.Save inserts or updates the client.
func (c *Client) Update() error {
	if c.ID == 0 {
		return fmt.Errorf("Cannot update client if ID = 0")
	}

	return c.service.Update(c)
}

// Client.Delete deletes the client.
func (c *Client) Delete() error {
	if c.ID == 0 {
		return fmt.Errorf("cannot delete client id = 0")
	}

	return c.service.Delete(c.ID)
}
