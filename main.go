package main

import (
	"errors"
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

	port := 3000
	startWebServer(port)
	err := startWebServer(port)
	fmt.Println(err)

}

func startWebServer(port int) error{
	
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)

	return errors.New("Something went wrong")
}