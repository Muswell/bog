package bog

import (
	"testing"
)

func TestClientCrud(t *testing.T) {
	c := NewClient()
	c.Name = "Test Client"
	c.Address1 = "123 Fake Street"
	c.Address2 = "Suite 100"
	c.City = "Branson"
	c.State = "MI"
	c.Zip = "12345"

	err := c.Insert()
	if err != nil {
		t.Errorf("Could not insert client: %v", err)
	}

	if c.ID == 0 {
		t.Error("Client insert failed: ID = 0")
	}

	c2, err := GetClient(c.ID)

	if err != nil {
		t.Errorf("Failed to load client: %v", err)
	}

	if c2.ID != c.ID {
		t.Errorf("Client ID does not match expected %d got %d", c.ID, c2.ID)
	}

	if c2.Name != c.Name {
		t.Errorf("Client Name does not match expected %s got %s", c.Name, c2.Name)
	}

	if c2.Address1 != c.Address1 {
		t.Errorf("Client Address1 does not match expected %s got %s", c.Address1, c2.Address1)
	}

	if c2.Address2 != c.Address2 {
		t.Errorf("Client Address2 does not match expected %s got %s", c.Address2, c2.Address2)
	}

	if c2.City != c.City {
		t.Errorf("Client City does not match expected %s got %s", c.City, c2.City)
	}

	if c2.State != c.State {
		t.Errorf("Client State does not match expected %s got %s", c.State, c2.State)
	}

	if c2.Zip != c.Zip {
		t.Errorf("Client Zip does not match expected %s got %s", c.Zip, c2.Zip)
	}

	c.Name = "Updated"

	err = c.Update()

	if err != nil {
		t.Errorf("Failed to update client: %v", err)
	}

	c2, err = GetClient(c.ID)

	if err != nil {
		t.Errorf("Failed to load client: %v", err)
	}

	if c2.Name != "Updated" {
		t.Errorf("Updated client name does not match expected Updated got %s", c2.Name)
	}

	err = c.Delete()

	if err != nil {
		t.Errorf("Failed to delete client: %v", err)
	}

	_, err = GetClient(c2.ID)

	if err == nil {
		t.Errorf("Delete client failed expected error when loading")
	}
}

func TestGetAllClients(t *testing.T) {
	c1 := NewClient()
	c1.Name = "Test Client 1"
	err := c1.Insert()
	if err != nil {
		t.Errorf("Could not insert client %v", err)
	}

	c2 := NewClient()
	c2.Name = "Test Client 2"
	err = c2.Insert()
	if err != nil {
		t.Errorf("Could not insert client %v", err)
	}

	clients, err := GetAllClients()

	if err != nil {
		t.Errorf("Could not load clients %v", err)
	}

	if len(clients) != 2 {
		t.Errorf("LoadAllClients failed count expected: 2 count received: %d", len(clients))
	}

	err = c1.Delete()
	if err != nil {
		t.Errorf("Could not delete client %v", err)
	}
	err = c2.Delete()
	if err != nil {
		t.Errorf("Could not delete client %v", err)
	}
}
