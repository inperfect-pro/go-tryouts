package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("contact added: %v\n", name)
}

func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("--- Listing Contacts ---")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone %s\n",
			i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("")
}

func main() {

	addContact("Alice", "alice@example.com", "111-2222")
	addContact("john", "john@example.com", "111-22232")
	addContact("edward", "edward@example.com", "111-22522")
	addContact("rene", "rene@example.com", "111-22622")

	ListContacts()

	bob := findContact("bob")
	if bob == nil {
		fmt.Println("No bob found.")
	} else {
		fmt.Println("Found bob: ", bob)
	}
}
