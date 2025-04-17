package main

import (
	"fmt"
	"sync"
	"time"
)

// ### Key Points to remember
// 1. sync.Pool is a concurrent safe pool of objects that can be reused to reduce memory allocations.
// 2. It is useful for temporary objects that are expensive to allocate and deallocate frequently.
// 3. The New function is called to create a new object when the pool is empty. It should return an interface{} type.
// 4. The Get method retrieves an object from the pool, and the Put method returns an object back to the pool.
// 5. The pool is not thread-safe, so you need to use a mutex or other synchronization mechanism if you want to share the pool across multiple goroutines.
// 6. The pool is not a replacement for a channel, but it can be used in conjunction with channels to manage object lifetimes.
// 7. The pool is not a garbage collector, so you need to be careful about memory leaks and object lifetimes.
// 8. The pool is not a replacement for a cache, but it can be used to cache objects that are expensive to create and destroy.

// note:
// - Great for short-lived objects that are expensive to allocate and deallocate frequently.
// - Objects in the pool may be garbage collected if they are not used for a while, so you should not rely on them being available at all times.

func main() {
	allocationCnt := 0
	// create a sync.Pool with a New function that initializes the objects.
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Print(".")
			allocationCnt++
			return make([]byte, 1024) // 1KB slice
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			obj := pool.Get().([]byte)
			fmt.Print("_")
			// simulate some work with the object
			time.Sleep(100 * time.Millisecond)
			pool.Put(obj)
			wg.Done()
		}()
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println("\n Total allocations:", allocationCnt)
}
