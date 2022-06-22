package main

import (
	"fmt"
	"time"
)

func main() {
	print_notes()
	// step1()
	print("-----------", "\n")
	// step2()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func step1() {
	greeter("Hello")
	greeter("World")
}

func step2() {
	go greeter("Hello")
	greeter("World")
}

func print_notes() {
	concurrency()
	parallelism()
	goroutines()
}

func concurrency() {
	println(`Concurrency :→
    → Get involved in doing the multiple tasks but not doing all the task at the same time
	→ In concurrency only one task at a time is done.
	→ At one instance of time you are doing only one task.
`)
}

func parallelism() {
	println(`Parallelism :→
    → Get involved in doing the multiple tasks but doing all the task at the same time
	→ In parallelism more than one task at a time is done.
	→ At one instance of time you are doing only more than one task.
`)
}

func goroutines() {
	println(`Goroutines:→
	   Thread                                                        Goroutines
	Managed by OS                                                 Managed by Go runtime
	Fixed Stack - 1MB (min)                                       Flexible stack - 2KB (min)
	`)
}
