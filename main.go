package main

import (
	"fmt"

	"github.com/KelvinWanyama/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Love",
		LastName: "Murumba",
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
	for{
		if x == 5 {
			break
		}
		println(x)
		x++
	}

}
