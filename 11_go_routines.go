//example code with go routines

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		//generate random sleep interval between 0 and 250 milliseconds
		d := time.Duration(rand.Intn(251)) * time.Millisecond
		time.Sleep(d)
	}
}

func main() {
	// synchronous function call
	f("direct")

	// asynchronous function call
	go f("goroutine 1")
	go f("goroutine 2")
	go f("goroutine 3")

	// anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// wait for goroutine to finish
	time.Sleep(time.Second)
	fmt.Println("done")
}
