package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Salary    int
	Position  string
	IsActive  bool
	JoinedAt  time.Time
}

// value receiver
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// pointer receiver
func (e *Employee) Deactivate() {
	e.IsActive = false
}

// pointer receiver
func deactivate(e *Employee) {
	e.IsActive = false
}

func main() {

	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Salary:    1000,
		Position:  "Manager",
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Println(jane.FullName())

	jane.Deactivate()
	fmt.Println(jane)

	deactivate(&jane)
	fmt.Println(jane)
}
