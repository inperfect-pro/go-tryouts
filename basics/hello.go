package main

import "fmt"

func main() {

	var greeting string = "Hello, world!"
	fmt.Println("Hello, World!")
	fmt.Println(1 + 1)
	fmt.Println(3.14)
	fmt.Println(true, false)
	fmt.Printf("%+v\n", []int{1, 2, 3})
	fmt.Printf("%+v\n", greeting)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	email := "test@test.com"
	fmt.Println(email)
}
