package main

import (
    "time"
    "fmt"
)

func main() {
    timeChan := make(chan time.Time)
    fmt.Println("Start at ", time.Now())
    go func() {
        time.Sleep(2*time.Second)
        timeChan <- time.Now()
    }()
    fmt.Println("Stop  at ", <-timeChan)
}