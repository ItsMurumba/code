package main

import "fmt"

func main() {
	fmt.Println("Hello from a module, Gophers!")

	/*
	Working with Arrays
	*/
	arr := [3]int{1,2,3}
	fmt.Println(arr)

	/*
	Working with Slices
	It is not a fixed sized entity as an array
	*/
	slice := []int{1,2,3}
	fmt.Println(slice)

	slice = append(slice,4,42,27)
	fmt.Println(slice)

	//Slice allows select given elements in an array to work with
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2,s3,s4)

	/*
	Working with Maps
	Maps are not read only
	*/
	m := map[string]int{"foo":42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"]=11
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	/*
	Working with Structs
	We can associate any type of data together
	*/
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	var u user
	u.ID = 1
	u.FirstName = "Kelvin"
	u.LastName = "Murumba"
	fmt.Println(u)

	//Single line initialization of user
	u2 := user{ID: 1,
		FirstName: "Mercy",
		LastName: "Nomzano",
	}
	fmt.Println(u2)




}