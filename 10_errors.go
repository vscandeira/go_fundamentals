//example code with errors using the MyError struct like documentation

package main

import (
	"fmt"
)

// MyError struct
type MyError struct {
	Msg string
}

// Error method for MyError struct
func (e *MyError) Error() string {
	return e.Msg
}

// Sqrt function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &MyError{"Negative number"}
	}
	return f * f, nil
}

func main() {
	// Sqrt function
	// it will return an error if the number is negative
	if f, err := Sqrt(-1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(f)
	}
}
