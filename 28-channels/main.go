// ref - https://youtu.be/i5HEE5Ikv3w
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	// // fatal error: all goroutines are asleep - deadlock!
	// myCh <- 5
	// fmt.Println(<-myCh)
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine
	wg.Add(2)
	// Recieve only
	go func(ch chan int, wg *sync.WaitGroup) {
		// close(myCh) // now using this line will through an error
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// Send Only
	go func(ch chan int, wg *sync.WaitGroup) {

		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
