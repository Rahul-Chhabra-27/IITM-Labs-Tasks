package main;

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel for communication b/w subroutines.
	channel := make(chan string);

	for _, link := range links {
		// call makeRequest function and pass the individual 
		// link and channel as parameters and use concurrency.
	}

	/*
	Print the messages sent by child subroutines...
	*/
}

func makeRequest(link string, channel chan string) {
	// complete the function.
}