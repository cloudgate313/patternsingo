package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Microsecond,
	)
}

func Producer(ch chan<- int, name string) {
	for {
		Sleep()

		n := rand.Intn(300)

		fmt.Println("Channel %s -> %d\n", n)

		ch <- n
	}
}

func Consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC = <-n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go Producer(chA, "A")
	go Producer(chB, "B")
	go Producer(chC, "C")

	fanIn(chA, chB, chC)
}
