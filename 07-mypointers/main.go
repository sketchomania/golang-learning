package main

import "fmt"

func main() {
	fmt.Println("Welcome to short intro of pointers in go")

	// var three int = 4
	// var two int
	// var ptr *int
	// var ptr2 *string
	// fmt.Println("Values of pointers are: ", three, ptr, ptr2, two)

	myNumber := 34

	var ptrnum = &myNumber
	// Here a pointer is created which is referencing a memory address
	// *ptr = actual value stored
	fmt.Println("Value of actual pointer is: ", ptrnum)
	fmt.Println("Value of actual pointer is: ", *ptrnum)

	*ptrnum = *ptrnum * 2
	fmt.Println("Updated value of ptrnum is: ", ptrnum)
	fmt.Println("Updated value of myNumber is: ", myNumber, *ptrnum)
}
