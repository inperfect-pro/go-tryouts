package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("--- while style ---")

	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("--- potentially infinite loop ---")
	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter == 10 {
			break
		}
	}

	fmt.Println("--- skipping ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--- array ---")
	items := [3]string{"Go", "Python", "Java"}
	for index, value := range items {
		fmt.Println(index, value)
	}
}
