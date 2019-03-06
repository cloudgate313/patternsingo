package main

import (
	"fmt"
	"math/rand"
	"time"
)

func EchoWorker(in chan int, out chan int) {
	for {
		n := <-in

		time.Sleep(
			time.Duration(rand.Intn(3000)) * time.Millisecond,
		)

		out <- n
	}
}

func Producer(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n", i)
		ch <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	// Worker pool of 20///
	for i := 0; i < 20; i++ {
		go EchoWorker(in, out)
	}

	go Producer(in)

	for n := range out {
		fmt.Printf("<- Receive Job: %d\n", n)
	}
}
