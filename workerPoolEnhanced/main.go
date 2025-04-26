package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Site represents a website to be crawled
type Site struct {
	URL string
}

// Result contains the outcome of a crawling operation
type Result struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}

// crawl processes URLs from the jobs channel and sends results to the results channel
// It terminates when the jobs channel is closed or the context is canceled
func crawl(ctx context.Context, wID int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a custom HTTP client with reasonable timeouts
	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     30 * time.Second,
		},
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d: shutting down due to context cancellation: %v", wID, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				log.Printf("Worker %d: jobs channel closed, exiting", wID)
				return
			}

			log.Printf("Worker %d: processing %s", wID, job.URL)
			startTime := time.Now()

			// Create request with context to enable cancellation
			req, err := http.NewRequestWithContext(ctx, "GET", job.URL, nil)
			if err != nil {
				log.Printf("Worker %d: error creating request for %s: %v", wID, job.URL, err)
				results <- Result{
					URL: job.URL,
					Err: err,
				}
				continue
			}

			// Add a user agent header
			req.Header.Set("User-Agent", "GolangCrawler/1.0")

			resp, err := client.Do(req)
			duration := time.Since(startTime)

			if err != nil {
				log.Printf("Worker %d: error fetching %s: %v", wID, job.URL, err)
				results <- Result{
					URL:      job.URL,
					Err:      err,
					Duration: duration,
				}
				continue
			}

			// Always close the body to prevent resource leaks
			defer resp.Body.Close()

			results <- Result{
				URL:        job.URL,
				StatusCode: resp.StatusCode,
				Duration:   duration,
			}
			log.Printf("Worker %d: completed %s with status %d in %v", wID, job.URL, resp.StatusCode, duration)
		}
	}
}

func main() {
	fmt.Println("Advanced Worker Pool Example")

	// Create buffered channels for jobs and results
	jobs := make(chan Site, 10)
	results := make(chan Result, 10)

	// Create a context that can be canceled
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Printf("Received signal %v, initiating shutdown", sig)
		cancel()
	}()

	var wg sync.WaitGroup

	// Adjust worker count based on workload or system resources
	numWorkers := 5
	log.Printf("Starting %d workers", numWorkers)

	// Start worker pool
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
		{URL: "https://httpbin.org/delay/2"},
		{URL: "https://golang.org"},
		{URL: "https://example.com"},
	}

	// Feed jobs to the worker pool with context awareness
	go func() {
		for _, site := range sites {
			select {
			case <-ctx.Done():
				log.Println("Job producer: context canceled, stopping job submission")
				close(jobs)
				return
			case jobs <- site:
				log.Printf("Queued job: %s", site.URL)
			}
		}
		log.Println("All jobs queued, closing jobs channel")
		close(jobs)
	}()

	// Process results with proper termination
	resultsComplete := make(chan struct{})
	go func() {
		// Track statistics
		successCount := 0
		errorCount := 0
		totalTime := time.Duration(0)

		for res := range results {
			if res.Err != nil {
				fmt.Printf("❌ Error fetching %s: %v\n", res.URL, res.Err)
				errorCount++
			} else {
				statusSymbol := "✅"
				if res.StatusCode >= 400 {
					statusSymbol = "⚠️"
				}
				fmt.Printf("%s %s - Status: %d (%.2fs)\n",
					statusSymbol, res.URL, res.StatusCode, res.Duration.Seconds())
				successCount++
				totalTime += res.Duration
			}
		}

		// Print summary statistics
		fmt.Println("\n--- Crawling Summary ---")
		fmt.Printf("Total URLs: %d\n", successCount+errorCount)
		fmt.Printf("Successful: %d\n", successCount)
		fmt.Printf("Failed: %d\n", errorCount)
		if successCount > 0 {
			fmt.Printf("Average response time: %.2fs\n", totalTime.Seconds()/float64(successCount))
		}
		close(resultsComplete)
	}()

	// Wait for all workers to finish
	wg.Wait()

	// All workers have finished, close the results channel to signal result processor
	log.Println("All workers finished, closing results channel")
	close(results)

	// Wait for result processing to complete
	<-resultsComplete

	// Clean up
	cancel()
	fmt.Println("\nProgram completed successfully")
}

// Key Improvements:
// Better documentation - Added comments to explain code functionality
// Custom HTTP client - Added timeouts and connection pooling
// Improved error handling - All errors are properly captured and reported
// Performance tracking - Added duration measurement for each request
// Graceful shutdown - Added signal handling for clean termination
// Context propagation - Better handling of context cancellation
// Resource cleanup - Proper closing of response bodies and channels
// Statistics reporting - Added summary of crawling results
// Visual feedback - Added status symbols to easily identify successes/failures
// Improved logging - More descriptive log messages
// These changes make the code more robust, maintainable, and provide better observability into its operation.
