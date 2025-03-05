package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ping(c chan string) {
	for i := 0; i <= 10; i++ {
		randNum := rand.Intn(100)
		c <- "ping - " + strconv.Itoa(randNum)
		fmt.Println("function call ping - ", randNum)
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	var c chan string = make(chan string)
	go ping(c)
	go print(c)

	var input string
	fmt.Scanln(&input)
}
