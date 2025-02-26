// simple for loop
package main

import (
	"fmt"
)

func main() {
	vet := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(vet); i++ {
		fmt.Println(vet[i])
	}
}
