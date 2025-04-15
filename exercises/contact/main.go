package contact

import "errors"

type Contact struct {
	Name  string
	Email string
	Phone string
}

type Contacts []Contact

func (c *Contacts) AddContact(contactToAdd Contact) error {
	var found bool = false
	var newContacts []Contact

	for _, contact := range *c {
		if contact.Name == contactToAdd.Name {
			found = true
			break
		}
	}

	if found {
		return errors.New("contact already exists")
	}

	newContacts = append(newContacts, contactToAdd)

	*c = newContacts

	return nil
}

func (c *Contacts) RemoveContactByName(name string) bool {
	var found bool = false
	var newContacts []Contact

	for _, contact := range *c {
		if contact.Name != name {
			newContacts = append(newContacts, contact)
			continue
		}

		found = true
	}

	*c = newContacts

	return found
}

func (c *Contacts) UpdateContactByName(name string, newContact Contact) bool {
	var found bool = false
	var newContacts []Contact

	for _, contact := range *c {
		if contact.Name != name {
			newContacts = append(newContacts, contact)
			continue
		}

		found = true
		newContacts = append(newContacts, newContact)
	}

	*c = newContacts

	return found
}
