package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {

	F := boilingF
	C := (F - 32) * 5 / 9
	fmt.Printf("Boiling temperature = %g°F or %g°C\n", F, C)
	fmt.Printf("Type of boilingF: %T\n", boilingF)
	fmt.Printf("Type of F: %T\n", F)
	fmt.Printf("Type of C: %T\n", C)
}
