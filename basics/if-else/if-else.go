package main

import "fmt"

func main() {

	tmp := 25
	if tmp > 30 {
		fmt.Println("greater than 30")
	} else if tmp <= 30 && tmp > 20 {
		fmt.Println("lower than 30, greater than 20")
	}

	userAccess := map[string]bool{
		"Jane": true,
		"John": false,
	}

	if hasAccess, ok := userAccess["John"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("access not granted")
	}
}
