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

func Sender(ch chan<- int) {
	for {
		Sleep()
		n := rand.Intn(100)
		fmt.Printf(" -> %d\n", n)
		ch <- n
	}
}

func Receiver(ch <-chan int, name string) {
	for n := range ch {
		go func(num int) {
			fmt.Printf("%s <- %d\n", name, num)
		}(n)
	}
}

func Fanout(chA <-chan int, chB chan<- int, chC chan<- int) {
	for n := range chA {
		if n > 50 {
			chB <- n
		} else {
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go Sender(chA)
	go Receiver(chB, "chB")
	go Receiver(chC, "chC")

	Fanout(chA, chB, chC)

}
