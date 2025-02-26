package main

import "fmt"

func main() {
	var vet = [5]int{1, 2, 3, 4, 5}
	var vet2 [5]int
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento do vetor 1: %d\n", vet[i])
		fmt.Printf("Elemento do vetor 2: %d\n\n", vet2[i])
	}
}
