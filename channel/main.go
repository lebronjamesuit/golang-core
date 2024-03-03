package main

func main() {
	links := []string{
		"http:google.com",
		"http:youtube.com",
		"http:gov.co.uk",
		"http://www.amazon.co.uk",
	}

	for _, link := range links {
		print("hello ", link)
	}

}
