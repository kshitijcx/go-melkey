package main

import (
	"fmt"
)

func basics() {
	// fmt.Println("hello")
	//variables
	var name string = "kshitij"
	fmt.Printf("my name: %s %d\n", name, 1000)
	age := 22
	fmt.Printf("%d\n", age)
	var city string
	city = "Dharamshala"
	fmt.Printf("%s ", city)
	var country, continent string = "India", "Asia"
	fmt.Println(country, continent)
	var (
		isEmployed bool   = true
		salary     int    = 20000
		position   string = "intern"
	)
	fmt.Println(isEmployed, salary, position)

	//default values
	var num int
	var st string
	var fl float64 //%f
	var bo bool    //%t
	fmt.Println(num, st, fl, bo)

	//constants
	const pi = 3.14       //can be unused
	const number int = 32 //can be types and untyped
	const (
		mon  = 1
		tues = 2
		wed  = 3
	)

	//enum - no inbuilt enum just a workaround
	const (
		Jan int = iota + 1
		// Jan = iota + 1
		Feb
		Mar
		Apr
	)

	//functions
	//add->private (cant be exported), Add->public(can be exported)
	fmt.Println(add(2, 3))
	sum, prod := sumandprod(2, 6) //sum,_ := ... to ommit other value
	fmt.Println(sum, prod)
}

// single returns
func add(num1 int, num2 int) int {
	return num1 + num2
}

// multiple returns
func sumandprod(a, b int) (int, int) {
	return a + b, a * b
}
