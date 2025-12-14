package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()

	return out
}

// Fan‑out and fan‑in pattern
// Why fai-in and fan-out
// fan-out: distribution of work
// Starting multiple goroutines to distribute work.
// fan-in: combining results
// Merging results from multiple goroutines into a channel.
// When to use fan-out and fan-in?
// Independence: stage should be independent of previous stage in terms of data.
// Computational intensity: stage should be computationally expensive to time consuming.
// Considerations
// Ordering: order of results may not be guranteed.
// Error handling: handle error appropriately to prevent unexpected behaviour.

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 3, 0, 5, 7)

	c1 := sq(done, in)
	c2 := sq(done, in)

	for v := range merge(done, c1, c2) {
		fmt.Println(v)
	}

}
