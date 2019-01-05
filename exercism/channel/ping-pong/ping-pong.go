package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	msg := ""
	for i := 0; ; i++ {
		msg = "ping " + fmt.Sprintf("%d", i)
		fmt.Println(msg)
		c <- msg
	}
}

func ponger(c chan string) {
	for {
		c <- "\tPONG " + <-c
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//c := make(chan string)
	c := make(chan string, 2)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
