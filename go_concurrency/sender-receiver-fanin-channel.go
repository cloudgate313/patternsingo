package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond,
	)
}

func Sender(ch chan<- int, name string) {
	for {
		Sleep()
		n := rand.Intn(100)
		fmt.Printf("%s<- %d\n", name, n)
		ch <- n
	}
}

func Receiver(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

func FanIn(chA <-chan int, chB <-chan int, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go Sender(chA, "chA")
	go Sender(chB, "chB")

	go Receiver(chC)

	// Why does this need to be after the goroutines?
	FanIn(chA, chB, chC)

}
