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

	fmt.Printf("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.Position)
	fmt.Println(jane.FirstName)
	fmt.Println(jane.JoinedAt)

	janePtr := &jane
	fmt.Println(janePtr)
	janePtr.FirstName = "JaneChanged"
	fmt.Println(janePtr)
	jane.LastName = "DoeChng"
	fmt.Println(jane)
}
