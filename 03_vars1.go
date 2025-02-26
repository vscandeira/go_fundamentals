package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {

	F := boilingF
	C := (F - 32) * 5 / 9
	fmt.Printf("Boiling temperature = %g°F or %g°C\n", F, C)
}
