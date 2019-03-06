package main

import "fmt"

func Sum(args ...int) (result int) {
	for _, num := range args {
		result += num
	}
	return result
}

func main() {
	fmt.Printf("Sum everything %d\n", Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 300))
}
