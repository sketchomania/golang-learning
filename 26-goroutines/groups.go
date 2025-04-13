package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer

// defining mutex
// mutex provide a lock over the memory
// lock till the one goroutine is working and till the goroutine is writing anything inside that mutex does not allow anybody to use this memory
var mut sync.Mutex // pointer

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://amazon.in",
		"https://github.com",
	}

	for _, web := range websitelist {
		// 1
		// getStatusCode(web)
		// 2 (waiting for all goroutines to come back)
		// go getStatusCode(web)
		// 3 (using waitgroup)
		go getStatusCode(web)
		wg.Add(1)

	}

	// this should be called at the end of the method
	// because, this is only responsible for not allowing tha main method to finish
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// our responsbility to call .Done()
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS error in endpoint: ", endpoint, err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
