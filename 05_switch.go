//simple go script to read a number from the user input and print if it is even, odd or zero

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	switch {
	case number == 0:
		fmt.Println("The number is zero")
	case number%2 == 0:
		fmt.Println("The number is even")
	default:
		fmt.Println("The number is odd")
	}
}
