package main

import "fmt"

const LoginToken string = "green val" // LoginToken is declared public -> "L"

func main() {
	// var username string = "sketch"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n", username)

	// var isLoggedIn bool = true
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T \n", smallVal)

	// var smallFloat float32 = 255.45798374987334
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values  and some aliases
	// var anotherVariable int
	// fmt.Println(anotherVariable)
	// fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// var anotherString string
	// fmt.Println(anotherString)
	// fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type
	var website = "github.com/sketchomania"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
