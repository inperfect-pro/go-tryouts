package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"
	fmt.Println("Today is:", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday":
		fmt.Println("Work days")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println(t)
		}
	}

	checkType(true)
}
