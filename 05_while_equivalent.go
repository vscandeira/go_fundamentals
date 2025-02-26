//simple script to iterate through a list of 7 numbers with a while loop and print if each number is even, odd or zero

package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 0
	for i < len(numbers) {
		number := numbers[i]
		switch {
		case number == 0:
			fmt.Printf("The number %d is zero\n", number)
		case number%2 == 0:
			fmt.Printf("The number %d is even\n", number)
		default:
			fmt.Printf("The number %d is odd\n", number)
		}
		i++
	}
}
