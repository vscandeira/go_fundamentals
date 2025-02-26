package main

import (
	"fmt"
)

func main() {
	v1 := 42
	v2 := 3.142
	v3 := 0.867 + 0.5i

	res12 := float64(v1) + v2
	fmt.Printf("res12 (int to float): %v, type: %T\n", res12, res12)

	res21 := v1 + int(v2)
	fmt.Printf("res21 (float to int): %v, type: %T\n", res21, res21)

	fmt.Printf("v3: %v, type: %T\n", v3, v3)

	// convert v3 to float64 format and add it do v2
	res23 := v2 + real(v3)
	fmt.Printf("res23 (complex to float): %v, type: %T\n", res23, res23)
}
