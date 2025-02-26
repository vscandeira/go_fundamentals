package main

import (
	"fmt"
)

func main() {
	vet := []int{10, 21, 32, 43, 54}
	slice := make([]int, 3)
	//prints a small vector with 3 elements all equals 0
	fmt.Println("Empty slice declared with 'make':", slice)
	fmt.Println("Entire vector:", vet)

	//new slice with the first 3 elements of the vector vet
	slice2 := vet[:3]
	fmt.Println("Slice of first 3:", slice2)

	//new slice with the last 2 elements of the vector vet
	slice3 := vet[3:]
	fmt.Println("Slice of last 2:", slice3)

	//appended slices 3 and 2, in that order to a new slice
	slice4 := append(slice3, slice2...)
	fmt.Println("Appended slices:", slice4)

	//example with copy directive
	slice5 := make([]int, 2)
	copy(slice5, vet)
	fmt.Println("Copied slice:", slice5)
}
