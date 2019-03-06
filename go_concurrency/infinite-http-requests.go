package main

import (
	"fmt"
	"net/http"
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

	for {
		go checkLink(<-ch, ch)
	}

}
