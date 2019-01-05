//A goroutine is a function that is capable of running concurrently with other functions.
//To create a goroutine we use the keyword go followed by a function invocation
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("f(n=", n, ") : i=", i)
		amt := time.Duration(rand.Intn(2000))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
