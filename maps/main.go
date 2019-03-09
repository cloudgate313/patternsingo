package main

import "fmt"

func mapTraverse(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex, "\n")
	}
}

func main() {

	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"black": "#000000",
		"white": "#FFFFFF",
	}

	mapTraverse(colors)

	fmt.Println(colors)
	delete(colors, "black")
	fmt.Println("Deleting black key/pair...")
	fmt.Println(colors)

}
