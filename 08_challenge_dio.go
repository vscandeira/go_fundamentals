//code to print on the screen numbers within an interval received by user input that are divisible by 3

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

	fmt.Println("\nNumbers divisible by 3 in the interval:")

	for i := start; i <= end; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
