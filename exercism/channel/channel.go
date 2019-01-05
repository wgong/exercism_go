//Channels provide a way for two goroutines to
//communicate with one another and synchronize their execution
package main

import (
	"fmt"
	"time"
)

func pinger(N int, c chan string) {
	for i := 0; i < N; i++ {
		c <- "ping-" + fmt.Sprintf("%d", i)
		time.Sleep(time.Second * 2)
	}
}

func ponger(c chan string) {

	msg := <-c
	fmt.Println(msg, " >> PONG")
}

func main() {
	nBuf := 20
	c := make(chan string, nBuf)

	go pinger(nBuf, c)
	for i := 0; i < nBuf; i++ {
		go ponger(c)
	}

	var input string
	fmt.Scanln(&input)
}
