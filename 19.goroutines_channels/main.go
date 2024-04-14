package main

import (
	"fmt"
	"net/http"
)

type website string

func main() {
	websites := []website{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, site := range websites {
		// The following call is sequential and will take some time i.e. wait till
		// first link is checked. We could have done this parallelly!
		go site.checkLink(c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c) // we cannot read only one value from here we need to expect all the values so that all go routines are completed
		// reading from channels is blocking
	}
}

func (w website) checkLink(c chan string) {
	_, err := http.Get(string(w))
	if err != nil {
		fmt.Println("Error: website must be down", err)
		c <- "Might be down I think"
		// os.Exit(1)
		return
	}
	fmt.Println(w, "is up!")
	c <- "Site is up!"
}
