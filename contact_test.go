package bog

import (
	"testing"
)

func TestContactCrud(t *testing.T) {
	client := NewClient()
	err := client.Insert()
	if err != nil {
		panic("Could not insert client")
	}
	defer client.Delete()

	contact := NewContact()
	contact.ClientID = client.ID
	contact.FirstName = "Randall"
	contact.LastName = "Carlson"
	contact.Email = "abc@test.com"
	contact.IsPrimary = true

	err = contact.Insert()
	if err != nil {
		t.Errorf("Could not insert contact: %v", err)
	}

	if contact.ID == 0 {
		t.Error("Contact insert failed: ID = 0")
	}

	contact2, err := GetContact(contact.ID)

	if err != nil {
		t.Errorf("Failed to load contact: %v", err)
	}

	if contact2.ID != contact.ID {
		t.Errorf("Contact ID mismatch, expected: %d got %d", contact.ID, contact2.ID)
	}

	if contact2.ClientID != contact.ClientID {
		t.Errorf("Contact ClientID mismatch, expected: %d got %d", contact.ClientID, contact2.ClientID)
	}

	if contact2.FirstName != contact.FirstName {
		t.Errorf("Contact FirstName mismatch, expected: %s got %s", contact.FirstName, contact2.FirstName)
	}

	if contact2.LastName != contact.LastName {
		t.Errorf("Contact LastName mismatch, expected: %s got %s", contact.LastName, contact2.LastName)
	}

	if contact2.Email != contact.Email {
		t.Errorf("Contact Email mismatch, expected: %s got %s", contact.Email, contact2.Email)
	}

	if contact2.IsPrimary != contact.IsPrimary {
		t.Errorf("Contact IsPrimary mismatch, expected: %t got %t", contact.IsPrimary, contact2.IsPrimary)
	}

	contact.FirstName = "John"
	err = contact.Update()
	if err != nil {
		t.Errorf("Could not update contact: %v", err)
	}

	contact2, err = GetContact(contact.ID)
	if err != nil {
		t.Errorf("Failed to load contact: %v", err)
	}

	if contact2.FirstName != "John" {
		t.Errorf("Updated contact name does not match expected John got %s", contact2.FirstName)
	}

	err = contact.Delete()

	if err != nil {
		t.Errorf("Failed to delete contact: %v", err)
	}

	_, err = GetContact(contact2.ID)

	if err == nil {
		t.Errorf("Delete contact failed. Expected error when loading")
	}
}

func TestGetAllContacts(t *testing.T) {
	client := NewClient()
	err := client.Insert()
	if err != nil {
		panic("Could not insert client")
	}
	defer client.Delete()

	client2 := NewClient()
	err = client2.Insert()
	if err != nil {
		panic("Could not insert client2")
	}
	defer client2.Delete()

	contact := NewContact()
	contact.ClientID = client.ID
	contact.FirstName = "John"

	err = contact.Insert()
	if err != nil {
		panic("Could not insert contact")
	}
	defer contact.Delete()

	contact2 := NewContact()
	contact2.ClientID = client.ID
	contact2.FirstName = "Billy"

	err = contact2.Insert()
	if err != nil {
		panic("Could not insert contact2")
	}
	defer contact2.Delete()

	contact3 := NewContact()
	contact3.ClientID = client2.ID
	contact3.FirstName = "Julie"

	err = contact3.Insert()
	if err != nil {
		panic("Could not insert contact3")
	}
	defer contact3.Delete()

	a, err := GetClientContacts(client.ID)
	if err != nil {
		t.Errorf("Could not get contact list %v", err)
	}

	if len(a) != 2 {
		t.Errorf("GetClientContacts failed expected count: 2 got count %d", len(a))
	}

	if a[0].ClientID != client.ID {
		t.Errorf("GetClientContacts returned contacts for incorect client id expected %d got %d", client.ID, a[0].ClientID)
	}

	if a[1].ClientID != client.ID {
		t.Errorf("GetClientContacts returned contacts for incorect client id expected %d got %d", client.ID, a[1].ClientID)
	}

	b, err := GetClientContacts(client2.ID)
	if err != nil {
		t.Errorf("Could not get contact list %v", err)
	}

	if len(b) != 1 {
		t.Errorf("GetClientContacts failed expected count: 1 got count %d", len(b))
	}

	if b[0].ClientID != client2.ID {
		t.Errorf("GetClientContacts returned contacts for incorect client id expected %d got %d", client2.ID, b[0].ClientID)
	}
}

func TestPrimaryContact(t *testing.T) {
	client := NewClient()
	err := client.Insert()
	if err != nil {
		panic("Could not insert client")
	}
	defer client.Delete()

	client2 := NewClient()
	err = client2.Insert()
	if err != nil {
		panic("Could not insert client2")
	}
	defer client2.Delete()

	contact := NewContact()
	contact.ClientID = client.ID
	contact.FirstName = "John"
	contact.IsPrimary = true

	err = contact.Insert()
	if err != nil {
		panic("Could not insert contact")
	}
	defer contact.Delete()

	contact2 := NewContact()
	contact2.ClientID = client.ID
	contact2.FirstName = "Billy"
	contact2.IsPrimary = false

	err = contact2.Insert()
	if err != nil {
		panic("Could not insert contact2")
	}
	defer contact2.Delete()

	contact3 := NewContact()
	contact3.ClientID = client2.ID
	contact3.FirstName = "Julie"
	contact3.IsPrimary = true

	err = contact3.Insert()
	if err != nil {
		panic("Could not insert contact3")
	}
	defer contact3.Delete()

	primary1, err := GetPrimaryContact(client.ID)
	if err != nil {
		panic("Could not load primary contact")
	}

	if primary1.ID != contact.ID {
		t.Errorf("Incorrect Primary Contact returned expected id %d got %d", contact.ID, primary1.ID)
	}

	primary2, err := GetPrimaryContact(client2.ID)
	if err != nil {
		panic("Could not load primary contact")
	}

	if primary2.ID != contact3.ID {
		t.Errorf("Incorrect Primary Contact returned expected id %d got %d", contact3.ID, primary1.ID)
	}

	contact2.IsPrimary = true
	contact2.Update()

	primary1, err = GetPrimaryContact(client.ID)
	if err != nil {
		panic("Could not load primary contact")
	}

	if primary1.ID != contact2.ID {
		t.Errorf("Incorrect Primary Contact returned expected id %d got %d", contact2.ID, primary1.ID)
	}

	contact, err = GetContact(contact.ID)

	if contact.IsPrimary {
		t.Errorf("Expected original primary to be overwritten")
	}
}
