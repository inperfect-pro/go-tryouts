package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong")
	}

	fmt.Println("executed without a panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	mightPanic(true)
}

func main() {

	//panic("oh no!")

	//mightPanic(true)

	recoverable()

}
