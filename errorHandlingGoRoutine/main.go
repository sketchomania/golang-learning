package main

import (
	"fmt"
	"net/http"
)

func checkIfExist(done <-chan struct{}, urls ...string) (<-chan *http.Response, <-chan error) {
	responsec := make(chan *http.Response)
	errc := make(chan error)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				return
			default:
				res, err := http.Get(url)
				if err != nil {
					errc <- err
					continue
				}
				responsec <- res
			}
		}
		close(responsec)
		close(errc)
	}()

	return responsec, errc
}

func main() {
	done := make(chan struct{})

	responsec, errc := checkIfExist(done, "https://google.com", "http://localhost:300")

	for i := 0; i < 2; i++ {
		select {
		case res := <-responsec:
			fmt.Println("Status:", res.Status)
		case err := <-errc:
			fmt.Println("Errors:", err)
		}
	}

	close(done)
}
