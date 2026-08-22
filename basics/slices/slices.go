package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	names = append(names, "Jane")
	fmt.Println(names)

	items := make([]int, 3, 5)
	fmt.Println(items)
	fmt.Printf("items: %+v, len: %d, cap: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	fmt.Printf("items: %+v, len: %d, cap: %d\n", items, len(items), cap(items))
	fmt.Printf("%+v", items[3:6])
}
