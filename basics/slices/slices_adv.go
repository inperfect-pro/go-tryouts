package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("--- Advanced Slicing Operations ---")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original: %v, len: %d, cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5]
	fmt.Printf("original: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	s2 := original[:4]
	fmt.Printf("original: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	s3 := original[6:]
	fmt.Printf("original: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	s4 := original[:]
	fmt.Printf("original: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	fmt.Println(slices.Contains(s4, 8))
	s4 = append(s4, 1000)
	fmt.Printf("s4 modified original: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
}
