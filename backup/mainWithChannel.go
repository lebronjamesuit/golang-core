package backup

import (
	"fmt"
	"net/http"
)

/* Questions
How many GoRoutines and Channel have been created in this function?

*/

func mainWithChannel() {

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
		go getLinkWithChannel(link, c)

	}
	fmt.Println("End Forloop before 4 GoRoutines return")

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func getLinkWithChannel(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil {
		c <- "Send this msg to our channel: Channel link might be down " + link
	} else {
		c <- "Send this msg to our channel:" + res.Status + link
	}

}
