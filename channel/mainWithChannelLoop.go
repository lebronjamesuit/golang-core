package main

import (
	"fmt"
	"net/http"
	"time"
)

/* Questions
How many GoRoutines and Channel have been created in this function?

*/

func mainWithChannelLoop() {

	links := []string{
		"http://cdl.co.uk",
		"http://gov.uk",
		"http://youtube.com",
		"http://amazon.co.uk",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		fmt.Println("Open a new GoRoutine for each time call getLinkWithChannel ")
		go getLinkWithChannelLoop(link, c)

	}
	fmt.Println("End Forloop before 4 GoRoutines return")

	// Method 1
	// emit from channel is: <-c   (return type is a link)
	// <-c = Wait for the channel to emit some value;
	/* for {
		time.Sleep(1 * time.Second)
		go getLinkWithChannelLoop(<-c, c)
	} */

	// Method 2
	// range c = Wait for the channel to emit some value; and assign the value to l
	for l := range c {
		time.Sleep(3 * time.Second) // this would add a pause in our main func and put it to sleep?
		go getLinkWithChannelLoop(l, c)
	}

}

func getLinkWithChannelLoop(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " it might be down")
		c <- link
		return
	}
	fmt.Println(link, "It's up")
	c <- link
}
