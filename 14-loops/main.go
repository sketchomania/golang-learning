package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	//  slice of days
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i], " ", i)
	// }

	// for index, day := range days {
	// 	fmt.Println("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Println("the value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto sketch
		}

		if rougueValue == 5 {
			// break
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
		//  ++rougueValue is not allowed in golang
	}

	// lbel
sketch:
	fmt.Println("Jumping at sketchomania.com")
}
