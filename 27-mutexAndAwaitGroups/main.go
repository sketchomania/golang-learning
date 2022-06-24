package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - ")

	wg := &sync.WaitGroup{}
	// for solving race issue we use mutex here
	// mutex will lock & unclock the memory
	// mut := &sync.Mutex{}
	// we can also use read write mutex
	// if you just use Mutex (not RWMutex) you don't need to add RLock. Read is a quick operation, compare to write
	mut := &sync.RWMutex{}

	// you should not lock the resourse directly
	var score = []int{0}

	wg.Add(4)
	// wg.Add(1)
	// IIFE
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go Routine One :")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go Routine Two :")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go Routine Three :")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go Routine Four :")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

// command used in this
// go run .
// go run --race .
// error - exit status 66
// go run --race .
