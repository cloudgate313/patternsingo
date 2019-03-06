package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "May be down...")
		ch <- "Site might be down"
		return
	}
	fmt.Println(link, "Link is up!")
	ch <- "Yep it's up"
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

	// put fmt.Println(<-ch) inside a for loop because recieving from a channel is a blocking function
	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}

}
