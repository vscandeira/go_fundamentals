package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ping(c chan string, limit int) {
	for i := 0; i <= limit; i++ {
		randNum := rand.Intn(100)
		if c != nil && i > 0 {
			msg := <-c
			fmt.Println(msg)
		}
		c <- "ping - " + strconv.Itoa(randNum)
	}
}

func pong(c chan string, limit int) {
	for i := 1; i <= limit; i++ {
		msg := <-c
		fmt.Println(msg)
		c <- strings.Replace(msg, "ping", "pong", -1)
	}
}

func main() {
	var c chan string = make(chan string)
	x := 10
	go ping(c, x)
	go pong(c, x)

	//var input string
	//fmt.Scanln(&input)
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
