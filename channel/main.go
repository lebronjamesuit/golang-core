package main

import (
	"fmt"
	"net/http"
)

func main() {
	/* links := []string{
		"http://google.com",
		"http://youtube.com",
		"http://gov.co.uk",
		"http://amazon.co.uk",
	}

	for _, link := range links {
		getLink(link)

	} */

	// mainWithChannel()

	mainWithChannelLoop()

}

func getLink(link string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link)
		fmt.Println(res.Status)
	}

}
