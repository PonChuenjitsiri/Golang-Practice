package main

import "fmt"

func main() {

	var name string = "pon"
	age := 19
	name1 := "nop"
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	_ = name1

	car,cost := "Audi" , 5000
	fmt.Println(car, cost)
	car,year := "BMW" , 2004
	_ = year

	var (
		salary float64
		firstName string
		gender bool
	)

	fmt.Println(salary, firstName, gender)

	var a,b,c int

	fmt.Println(a,b,c)

	
}
