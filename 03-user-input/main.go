package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || comma err syntax
	// in go we don't have try catch
	// you treat problems or errors are something like variable or true/false

	// _, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is: %T", input)
}
