package main

import "fmt"

type myStruct struct {
	num int
}

// FACT:- Go does not allow you to do pointer arithmetic
// Pointer arithmetics are bad and against core priciples of go
// if you want to use it any case, go provide a package unsafe for it
func main() {
	fmt.Println("Welcome to pointers in golang")

	// exOne()

	// nullPointer()

	// exThree()

	// exFour()

	exFive()
}

func exOne() {
	// 1.Normal way of working
	var a int = 42
	fmt.Printf("[a] value = %v, type = %T\n", a, a)

	// var b int = a
	// fmt.Printf("[b] value = %v, type = %T\n", b, b)

	// b = 21
	// fmt.Printf("[a] value = %v, type = %T\n", a, a)
	// fmt.Printf("[b] value = %v, type = %T\n", b, b)

	var c *int = &a
	fmt.Printf("[c] value = %v, type = %T\n", c, c)
	// de-referencing poiner :- *c
	fmt.Printf("[*c] value = %v, type = %T\n", *c, *c)
	// fmt.Printf("[&c] value = %v, type = %T\n", &c, &c)

	// deReferencing
	// a = 21
	// fmt.Printf("[a] value = %v, type = %T\n", a, a)
	// fmt.Printf("[*c] value = %v, type = %T\n\n", *c, *c)

	// deReferencing
	// *c = 11
	// fmt.Printf("[a] value = %v, type = %T\n", a, a)
	// fmt.Printf("[*c] value = %v, type = %T\n\n", *c, *c)
}

func nullPointer() {
	var p *int
	// fmt.Printf("[p] value = %v, type = %T\n\n", p, p)

	// this deReferencing will cause error
	// fmt.Printf("[*p] value = %v, type = %T\n\n", *p, *p)

	if p != nil {
		fmt.Printf("[*p] value = %v, type = %T\n\n", *p, *p)
	} else {
		fmt.Printf("Null pointer dereferenced...!")
	}
}

func exThree() {
	// we can have pointer of any data type
	var g float64 = 0.42
	fmt.Printf("[g] value = %v, type = %T\n\n", g, g)

	var h *float64 = &g
	fmt.Printf("[h] value = %v, type = %T\n", h, h)
	// de-referencing poiner :- *h
	fmt.Printf("[*h] value = %v, type = %T\n", *h, *h)
}

func exFour() {
	// we can have pointer of any data type
	var i myStruct = myStruct{6}
	fmt.Printf("[i] value = %v, type = %T\n\n", i, i)

	var j *myStruct = &i
	// fmt.Printf("[j] value = %v, type = %T\n", j, j)
	// %p is for printing the address
	fmt.Printf("[j] value = %p, type = %T\n", j, j)
	// de-referencing poiner :- *j
	fmt.Printf("[*j] value = %v, type = %T\n\n", *j, *j)

	var ms *myStruct
	ms = &myStruct{10}
	// %p is for printing the address
	fmt.Printf("[ms] value = %p, type = %T\n", ms, ms)
	fmt.Printf("[*ms] value = %v, type = %T\n", *ms, *ms)
}

func exFive() {
	// Allocation primitives which is called the new function
	// The new function in go alloctes memory for a variable of a particular type in the memory and it returns the address to that memory location
	var k myStruct = myStruct{1}
	fmt.Printf("[k] value = %v, type = %T\n\n", k, k)

	var l *int = new(int)
	// new(int) is allocating memory for a variable of type int in the pc's memory and then it is returning the address of the newly allocted variable
	// here we storing that address in variable: " l "
	*l = 13
	fmt.Printf("[l] value = %p, type = %T\n", l, l)
	// %p is for printing the address
	fmt.Printf("[*l] value = %v, type = %T\n\n", *l, *l)
}
