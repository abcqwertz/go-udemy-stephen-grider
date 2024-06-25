package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.golang.org",
		"https://www.amazon.com",
		"https://www.reddit.com",
	}

	c := make(chan string)

	for _, website := range websites {
		go checkLink(website, c)

	}
	for w := range c {
		go func(website string) {
			time.Sleep(5 * time.Second)
			checkLink(website, c)
		}(w)
	}

}

func checkLink(website string, c chan string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "is down")
		c <- website
		return
	}
	fmt.Println(website, "is up and running")
	c <- website
}
