package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func calculateArea(width, height float64) float64 {
	if width < 0 || height < 0 {
		fmt.Println("Error: width and height must be positive")
		return 0
	}
	return width * height
}

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	greet("Gopher")

	add(1, 2)

	area := calculateArea(10, 20)
	fmt.Println(area)

	fmt.Println(factorial(6))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Hello World")
}
