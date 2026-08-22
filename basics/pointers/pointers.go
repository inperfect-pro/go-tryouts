package main

import "fmt"

func main() {
	age := 10
	agePtr := &age
	fmt.Printf("age: %d\n", &age)
	fmt.Printf("age: %d\n", agePtr)

	modifyValue(age)
	fmt.Printf("age: %d\n", age)
	modifyPointer(&age)
	fmt.Printf("age: %d\n", age)

	grade := 50
	gradePtr := &grade
	fmt.Printf("grade address: %+v\n", gradePtr)
	fmt.Printf("gradePtr address: %+v\n", &gradePtr)
}

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modified val: %d\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 // dereferencing
}
