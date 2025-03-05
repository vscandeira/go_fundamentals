// example code to create a struct of a client of a book store and use it
package main

import (
	"fmt"
)

// Client struct
type Client struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// creating a new client
	client := Client{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// printing the client
	fmt.Println("Client:", client)

	// changing the client's address
	client.Address = "456 Elm St"

	// printing the client
	fmt.Println("Client:", client)
}
