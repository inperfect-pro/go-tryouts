package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"James": 32,
		"Dan":   60,
		"Alice": 40,
	}
	fmt.Println(studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Println("Alice's student grade is", alice)
	}

	name := "Bob"
	if _, ok := studentGrades[name]; ok {
		fmt.Printf("%s student grade is %d", name, studentGrades[name])
	}

	delete(studentGrades, "Alice")
	fmt.Println(studentGrades)

	//configs := make(map[string]int)
	var configs map[string]int
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Println("configs is nil")
	}
}
