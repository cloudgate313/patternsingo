package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	channel := make(chan string)

	wg.Add(1)
	go func() {
		channel <- "Go Channel"
		fmt.Println("go func running...")
		wg.Done()
	}()
	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
	wg.Wait()

}
