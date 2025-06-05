package main

import "fmt"

func main() {
	//arrays
	numbers := [5]int{10, 20, 30, 40, 50}
	numbers[3] = 1100
	fmt.Println("array:", numbers, len(numbers))
	fmt.Printf("array: %v", numbers)
	fmt.Println(numbers[len(numbers)-1])

	numbersAtInit := [...]int{10, 20, 30}
	matrix := [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(numbersAtInit, matrix)

	//array slicing
	allNumbers := numbers[:] //array slicing
	allNumbers2 := numbers[0:3]
	fmt.Println(allNumbers, allNumbers2)

	fruits := []string{"Apple", "Banana", "Mango"}
	fruits = append(fruits, "Guava", "Kiwi")
	moreFruits := []string{"Grapes", "Melon"}
	fruits = append(fruits, moreFruits...)
	fmt.Println(fruits)

	//iterate through list
	for index, value := range numbers{
		fmt.Println(index, value)
	}

	//maps
	capitalCities:= map[string]string{
		"USA":"Washington DC",
		"India":"New Delhi",
		"UK":"London",
	}
	fmt.Println(capitalCities["India"])
	capital, exits :=capitalCities["Germany"]
	if exits{
		fmt.Println(capital)
	}else{
		fmt.Println(exits,"Does not exist")
	}

	delete(capitalCities,"UK")
	fmt.Println(capitalCities)
}
