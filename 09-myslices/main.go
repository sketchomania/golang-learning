package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to intro on slices in golang")

	// declaration of slice
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	//   [1:]    [:3]
	fmt.Println(fruitList)

	highScores := make([]int, 5)

	highScores[0] = 234
	highScores[1] = 334
	highScores[2] = 834
	highScores[3] = 234
	highScores[4] = 944
	//highScores[5] = 444

	highScores = append(highScores, 231, 666, 777)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "rust"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
