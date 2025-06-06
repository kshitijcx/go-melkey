package main

import (
	"fmt"
)

// structs
type Person struct {
	Name string
	Age  int
}

func ds() {
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
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	//maps
	capitalCities := map[string]string{
		"USA":   "Washington DC",
		"India": "New Delhi",
		"UK":    "London",
	}
	fmt.Println(capitalCities["India"])
	capital, exits := capitalCities["Germany"]
	if exits {
		fmt.Println(capital)
	} else {
		fmt.Println(exits, "Does not exist")
	}

	delete(capitalCities, "UK")
	fmt.Println(capitalCities)

	//using structs
	person := Person{"John", 12} //Person{Name:"John", Age: 25}
	fmt.Println(person)          //%v or %+v this gives both key and values

	//anonymous strcut
	employee := struct {
		Name string
		Age  int
	}{"Raman", 25}
	fmt.Println(employee)

	//in go everything is passed by value and explicitely say by reference

	//nested structs
	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{"Marc", Address{"Sidhbari", "Dharamshala"}, "9021323434"}
	fmt.Println(contact)
	//If values are not passed then default values will be assigned like 0 for int

	//pointers
	fmt.Println(person.Name)
	// modifyPersion(&person)
	person.modifyPersion("Superman")
	fmt.Println(person.Name)

	x:=20
	ptr:=&x
	fmt.Printf("%d %p\n",x,ptr)
	*ptr=30
	fmt.Println(ptr,*ptr)
}

// func modifyPersion(person *Person){
// 	person.Name = "Ganesh"
// }
func (p *Person) modifyPersion(name string){ //exists only for Person struct (p *Person) -> Method Receiver
	p.Name = name
}