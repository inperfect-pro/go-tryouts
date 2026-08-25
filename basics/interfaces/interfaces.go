package main

import "fmt"

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func displayPerson(emp Person) {
	fmt.Println(emp.GetName())
}

func main() {

	joe := Employee{
		1,
		"Joe",
	}
	displayPerson(joe)

	jane := BusinessPerson{
		1,
		"Jane",
	}

	displayPerson(jane)
}
