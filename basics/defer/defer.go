package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred") // comes at the end
	fmt.Println("Function simpleDefer: End")
}

func lifoSimpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: first deferred")
	defer fmt.Println("Function simpleDefer: last deferred")
	fmt.Println("Function simpleDefer: End")
}

func main() {
	//simpleDefer()

	defer func() {
		fmt.Println("Function deferred before end of main()")
	}()

	lifoSimpleDefer()

}
