package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to learning goroutines(concurrency) in golang")
	// print_notes()
	// hello(2)
	// go hello(1) // Goroutine [ Concurrency ]
	// go hello(3) // Goroutine [ Concurrency ]
	// go hello(4) // Goroutine [ Concurrency ]
	// go hello(5) // Goroutine [ Concurrency ]
	// go hello(6) // Goroutine [ Concurrency ]
	// go hello(7) // Goroutine [ Concurrency ]

	msg := "hello"
	// annonimus function
	func() {
		fmt.Println("From anno. func: " + msg)
	}()
	// not ok
	// go func() {
	// 	fmt.Println("From anno. func: " + msg)
	// }()
	// kinda ok
	go func(m string) {
		fmt.Println("From anno. func: " + m)
	}(msg)
	msg = "world" // this will create a raise condition

	fmt.Println("Main ends")
	// goto sleep for cretain amount of time
	time.Sleep(100 * time.Millisecond)
}

func hello(grn int) {
	for i := 0; i < 3; i++ {
		fmt.Println(grn, ": Hello")
	}
}
func print_notes() {
	definition()
	benifits()
	threads()
	goroutines()
	advantages()
}

func definition() {
	fmt.Println(`"Concurrency" is the ability of different parts of units of a program,
	 algorithm, or problem to be executed out-of-order or in partial order,
	 without affecting the final outcome.`)
}
func benifits() {
	fmt.Println(` Benifits :-
	 This allow for parallel execution of the cocurrent units,
	 which can significantly improve overall speed of the execution in multi-processor and multi-core systems.
	 "Parallelism" is the result of concurrency.
	 Example: 
		(1+2) * (3+4) => 3 * (3+4) => 3 * 7 => 21	[3 step process]
		(1+2) * (3+4) => 3 * (3+4) => 3 * 7 => 21	[3 step process]
		(1+2) * (3+4) => 3 * 7 => 21	[2 step process]
	
	"This is only useful in case of multi-processors and multi-core system"
	`)
}
func threads() {
	fmt.Println(`Threads :-
	○ Concurrency achieved through threads.
	○ Most languages use OS threads (managed, scheduled etc by OS)
	○ This means that the threads have am individual function call stack dedicated to the
	  execution of the code of the thread
	○ As a result this Take Too many resources: 1 MB RAM, take time to setup
	○ Leads to threads pools as creation and destruction of threads is difficult
	`)
}
func goroutines() {
	fmt.Println(` Goroutines :-
	○ GO Threads are "Green Threads"(threads which are managed by user program, user library and so on...)
	○ Instead of massively heavy and resource intensive threads, Go creates an abstraction
	  of a thread called a "goroutine"
	○ Go runtime has a scheduler tha maps goroutines into OS threads for certain periods of 
	  time and keeps on assigning the goroutine some amounts of processing time on those CPU threads.
	○ We don't interact with low level threads, but we interact with the high level goroutines.
	`)
}
func advantages() {
	fmt.Println(` Advantages :-
	○ Goroutines can be started with very small stack spaces (not very resource intensive)
	○ As they can be reallocated very quickly they can be easily created or destroyed
	○ In a Go application we can have 1000s of goroutines and have not issues
	○ Threads in other languages with 1 MB overhead cannot scale in this manner
	`)
}
