//example code to demonstrate the use of pointers in Go
//it defines an integer x variable with value 5,
// then with a function receiving x as parameter initializes x to 0 and
// prints the value of x

package main

import "fmt"

// zero function to set the value of x to 0
func zero(x *int) {
	*x = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is still 5
}
