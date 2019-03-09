package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreating(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreating(eb)
	printGreating(sp)
}
