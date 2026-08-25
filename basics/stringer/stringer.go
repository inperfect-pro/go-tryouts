package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func (e BusinessPerson) String() string {
	return fmt.Sprintf("Name: %s, ID: %d", e.Name, e.ID)
}

func displayPerson(emp Person) {
	fmt.Println(emp.GetName())
}

func main() {

	jane := BusinessPerson{
		1,
		"Jane",
	}

	fmt.Println(jane)
}
