package main

import "fmt"

//A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.
func main() {
	fmt.Println("Defer in golang")
	// execute a function at the end
	// LIFO  --> last in first out

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	fmt.Println("At myDefer()")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
