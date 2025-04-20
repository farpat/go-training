package contact

import "testing"

func TestItAddsContactProperly(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts

	// ACT
	contacts.AddContact(input)

	// ASSERT
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}

	if contacts[0] != input {
		t.Errorf("Expected contact %v, got %v", input, contacts[0])
	}
}

func TestItFailsWhenAddingExistingContact(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts

	// ACT
	contacts.AddContact(input)

	// ASSERT
	err := contacts.AddContact(input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestItRemovesContactByName(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts = Contacts{input}

	// ACT
	contacts.RemoveContactByName(input.Name)

	// ASSERT
	if len(contacts) != 0 {
		t.Errorf("Expected 0 contacts, got %d", len(contacts))
	}
}

func TestItDoesNothingIfWeTryToRemoveNonExistingContact(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts = Contacts{input}

	// ACT
	isRemoved := contacts.RemoveContactByName("Jane Doe")

	// ASSERT
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}
	if isRemoved {
		t.Errorf("Expected false, got true")
	}
}

func TestItUpdatesContactByName(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts = Contacts{input}

	// ACT
	isUpdated := contacts.UpdateContactByName(input.Name, Contact{Name: "John Doe updated", Email: "john.doe.updated@example.com", Phone: "0987654321"})

	// ASSERT
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}
	if !isUpdated {
		t.Errorf("Expected true, got false")
	}
	if !isEquals(contacts[0], Contact{Name: "John Doe updated", Email: "john.doe.updated@example.com", Phone: "0987654321"}) {
		t.Errorf("Expected contact %v, got %v", Contact{Name: "John Doe updated", Email: "john.doe.updated@example.com", Phone: "0987654321"}, contacts[0])
	}
}

func TestItDoesNothingIfWeTryToUpdateNonExistingContact(t *testing.T) {
	// ARRANGE
	input := Contact{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"}
	var contacts Contacts = Contacts{input}

	// ACT
	isUpdated := contacts.UpdateContactByName("Jane Doe", Contact{Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "0987654321"})

	// ASSERT
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}
	if isUpdated {
		t.Errorf("Expected false, got true")
	}
	if !isEquals(contacts[0], input) {
		t.Errorf("Expected contact %v, got %v", input, contacts[0])
	}
}

func isEquals(c1 Contact, c2 Contact) bool {
	return c1.Name == c2.Name && c1.Email == c2.Email && c1.Phone == c2.Phone
}
