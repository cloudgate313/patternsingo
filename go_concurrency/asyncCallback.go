package main

import (
	"fmt"
	"strings"
	"sync"
)

func asyncToUpper(words string, f func(string)) {
	go func() {
		f(strings.ToUpper(words))
	}()
}

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	asyncToUpper("Get Money", func(v string) {
		fmt.Printf("Callback: %s\n", v)
		wg.Done()
	})
	fmt.Println("Async function running...")
	wg.Wait()
}
