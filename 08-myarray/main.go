package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[2] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))
}
