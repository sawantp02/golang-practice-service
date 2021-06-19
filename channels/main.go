package main

import (
	"net/http"
	"fmt"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for _, link := range links {
		fmt.Println(link, <-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "might be down!"
		return
	}

	c <- "is up!"
}