//code to print on the screen numbers within an interval, received by user input,
// when an number is divisible by 3, print "Pin" instead of the number
// when an number is divisible by 5, print "Pan" instead of the number
// when an number is divisible by 3 and 5, print "PinPan" instead of the number
// when an number is not divisible by 3 or 5, print the number

package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Print("Enter the start of the interval: ")
	fmt.Scan(&start)

	fmt.Print("Enter the end of the interval: ")
	fmt.Scan(&end)

	fmt.Println("\nLet's play :)")

	for i := start; i <= end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
