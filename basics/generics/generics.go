package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T Number](numbers ...T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {

	grades := []int{90, 85}
	people := []string{"Jane", "John", "Mark"}
	fmt.Println(len(grades), len(people))

	fmt.Println(Sum(30.2, 1.5))

	v := Sum(10, 2, 3)
	fmt.Printf("%T\n", v)
}
