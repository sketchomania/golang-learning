package main

import "fmt"

// WE can't declare function inside a function in golang
func main() {
	fmt.Println("Functions in golang...")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 6, 3, 8, 9, 7)
	fmt.Println("Prp result is: ", proRes)
	fmt.Println("Prp Message is: ", myMessage)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
