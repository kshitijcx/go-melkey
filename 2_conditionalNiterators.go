package main

import "fmt"

func conditionalsNiterators() {
	//conditionals
	age := 30
	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Thursday"

	switch day {
	case "Monday":
		fmt.Println("First day")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Mid week")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("Weekend")
	}

	//for loops 
	for i := 0; i < 5; i++ {
		fmt.Println("this is i:", i)
	}

	//while loop -> break and continue
	counter := 0
	for counter < 3 {
		fmt.Println("counter: ", counter)
		counter++
	}

	//infinite loop 
	iter:=0
	for{
		if iter==5{
			fmt.Println("iter=",iter)
			break
		}
		iter++
	}

}
