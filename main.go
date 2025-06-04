package main

import "fmt"

func main() {
	//conditionals
	age := 30
	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day:="Thursday"

	switch day {
		case "Monday":
			fmt.Println("First day")
		case "Tuesday","Wednesday","Thursday":
			fmt.Println("Mid week")
		case "Friday":
			fmt.Println("TGIF")
		default:
			fmt.Println("Weekend")
	}
}
