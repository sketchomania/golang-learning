package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Site struct {
	URL string
}

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func crawl(ctx context.Context, wID int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: context cancelled\n", wID)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", wID)
				return
			}

			// for job := range jobs {
			log.Printf("Worker ID: %d started  %s", wID, job.URL)
			req, _ := http.NewRequestWithContext(ctx, "GET", job.URL, nil)
			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				log.Printf("Worker ID: %d error %s", wID, err)
				results <- Result{StatusCode: 0, Err: err}
				// continue
			} else {

				results <- Result{
					URL:        job.URL,
					StatusCode: resp.StatusCode,
				}
				resp.Body.Close()
				log.Printf("Worker ID: %d finished %s", wID, job.URL)
			}
		}
	}
}

func main() {
	fmt.Println("Worker Pool Example")

	jobs := make(chan Site, 10)
	results := make(chan Result, 10)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	numWorkers := 3
	// numJobs := 6

	// start worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go crawl(ctx, w, jobs, results, &wg)
	}

	sites := []Site{
		{URL: "https://www.google.com"},
		{URL: "https://www.example.com"},
		{URL: "https://www.github.com"},
		{URL: "https://www.stackoverflow.com"},
		{URL: "https://www.reddit.com"},
		{URL: "https://www.medium.com"},
		{URL: "https://www.linkedin.com"},
		{URL: "https://www.twitter.com"},
		{URL: "https://www.facebook.com"},
		{URL: "https://httpbin.org/get"},
		{URL: "https://httpbin.org/status/404"},
		{URL: "https://httpbin.org/status/500"},
		{URL: "https://httpbin.org/delay/2"}, // will test timeout behavior
		{URL: "https://golang.org"},
		{URL: "https://example.com"},
	}

	// Feed jobs to the worker pool
	go func() {
		for _, site := range sites {
			jobs <- Site{URL: site.URL}
		}
		close(jobs)
	}()

	// Read results
	go func() {
		for res := range results {
			if res.Err != nil {
				fmt.Printf("Error fetching %s: %v\n", res.URL, res.Err)
			} else {
				fmt.Printf("Fetched %s - StatusCode: %d\n", res.URL, res.StatusCode)
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(results)
	fmt.Println("All workers finished")
	fmt.Println("Main function finished")
	fmt.Println("Exiting program")
}
