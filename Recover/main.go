// Recover is a built in functions that regains control of a panicking goroutine.
// Recover is only useful inside defered functions.
// During normal execution, a call to recover will return nil and have no other effect.
// If the current goroutine id panicking a call to recover will capture the value given to panic and resume noraml execution.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main: start learning to recover in Go")
	// iCausePanic1()

	iCausePanic2()

	// iCausePanic3()

	fmt.Println("Main: Post Panic")
}

func iCausePanic1() {
	fmt.Println("about to panic")
	panic("func: cannot continue to execute")
	fmt.Println("i just panicked 😁")
}

func iCausePanic2() {
	defer tryToRecover()
	fmt.Println("about to panic")
	panic("iCausePanic2: func: cannot continue to execute")
	fmt.Println("i just panicked 😁")
	// execution will brought bavk to the panic function
}

func iCausePanic3() {
	defer tryToRecover()
	fmt.Println("about to panic")
	// panic("func: cannot continue to execute")
	fmt.Println("i just panicked 😁")
	// execution will brought bavk to the panic function
}

func checkNilError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func tryToRecover() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from Error: ", err)
		panic("tryToRecover: cannot recover from panic")
	} else {
		fmt.Println("Nothing to recover from: ", err)
	}
}
