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

	contact2, err := LoadContact(contact.ID)

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

	contact2, err = LoadContact(contact.ID)
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

	_, err = LoadClient(contact2.ID)

	if err == nil {
		t.Errorf("Delete contact failed. Expected error when loading")
	}
}
