package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "May be down...")
		ch <- link
		return
	}
	fmt.Println(link, "Link is up!")
	ch <- link
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.org",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}
