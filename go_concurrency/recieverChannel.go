package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "Go channel"
		println("Finish go routine")
	}(channel)
	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
}
