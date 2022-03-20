package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - sketchomania.com"

	file, err := os.Create("./mysketchgofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	// panic will shut down the execution of program and show the error
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("lenght is: ", length)
	defer file.Close()

	readFile("./mysketchgofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is: \n", databyte)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		// panic will shut down the execution of program and show the error
		panic(err)
	}
}
