// There are no exceptions in Go
// Exceptional situations are considered normal
//   Like: Opening a file that does not exist, is reasonable, not exceptional
// Errors are retuned tho handle such situations
// Sometimes a Go program cannot figure ou what to do and cannot continue to run.
// The program Panics in this case.
// panic will shut down the execution of program and show the error

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Panics in Go")
	// iCausePanic1()

	// iCausePanic2()

	// iCausePanic3()

	iCausePanic4()
}

func iCausePanic1() {
	a, b := 10, 0
	fmt.Println(a / b)
}

func iCausePanic2() {
	panic("Cannot continue the program 🤦‍♂️ ")
	fmt.Println("Post Panic")
}

func iCausePanic3() {
	fmt.Println("1 Start")
	defer fmt.Println("2 Deffered")
	panic("Cannot continue the program 🤦‍♂️ ")
	fmt.Println("Post Panic")
}

func iCausePanic4() {
	// s := "100"
	s := "xyz"
	n, err := strconv.ParseInt(s, 10, 64)
	checkNilError(err)
	fmt.Println(n)
}

func checkNilError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
