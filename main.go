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

	//returning multiple values
	port, err2 := startWebServer2(port)
	fmt.Println(port, err2)

}

func startWebServer(port int) error{
	
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)

	return errors.New("Something went wrong")
}

//returning multiple values
func startWebServer2(port int) (int,error){
	
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return port, nil
}