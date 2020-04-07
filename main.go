package main

import "fmt"

//Constant blocks
const (
	first = iota + 6
	second 
)
func main() {
	fmt.Println("Hello from a module, Gophers!")

	var lastName *string = new(string)
	*lastName = "Murumba"
	fmt.Println(*lastName)

	/*
	Working with Pointers
	*/
	firstName := "Kelvin"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Trish"
	fmt.Println(ptr, *ptr)

	/*
	Working with Constants
	Constants are declared at the initial time
	*/
	const pi = 3.14
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)

	fmt.Println(c + 1.2)

	/*
	Working with iota
	iota resets everytime it is in a new const block
	*/
	fmt.Println(first, second)



}