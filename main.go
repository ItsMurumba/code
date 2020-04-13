package main

import (
	"fmt"

	"github.com/KelvinWanyama/webservice/models"
)

// User :
type User struct {
	ID        int
	FirstName string
	LastName  string
}

//HTTPRequest :
type HTTPRequest struct{
	Method string
}

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Love",
		LastName:  "Murumba",
	}
	fmt.Println(u)

	/*
		Working with Loops
	*/
	var i int
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			continue
		}
		println("Continuing.....")

		if i == 4 {
			break
		}
	}

	//Implicit initalization within for loop scope
	for i := 0; i < 5; i++ {
		println(i)
	}
	// println(i) variable i will not be accessible here.

	//Infinite loop
	var x int
	for {
		if x == 5 {
			break
		}
		println(x)
		x++
	}

	/*
		Loops with collections: arrays, slices
		When using only value, use _ in the for statement
	*/
	slice := []int{1, 2, 3}
	for y, v := range slice {
		println(y, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for y, v := range wellKnownPorts {
		println(y, v)
	}

	//Working with the Panic Function
	println("Starting web server")

	//do important things
	//panic("Man down, I repeat man down, Someone call 911!")

	println("Web server started")

	//Working with if else
	u1 := User{
		ID:        1,
		FirstName: "Mercy",
		LastName:  "Nomzano",
	}
	u2 := User{
		ID:        2,
		FirstName: "Kelvin",
		LastName:  "Murumba",
	}

	if u1.ID == u2.ID {
		println("Same user")
	} else {
		println("Not same user")
	}


	/*
	working with switch statements
	Has no default case but can define a special default case.
	Has break implicitly implemented so no need to define it
	*/
	r := HTTPRequest{Method: "GET"}

	switch r.Method{
	case "GET":
		println("Get request")
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}


}
