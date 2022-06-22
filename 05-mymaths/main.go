package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4
	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// Random Number:-  1.[math/rand] 2.[crypto/rand]

	// [math/rand]
	// 1
	// fmt.Println("1. This will print same number every time", rand.Intn(6))
	// 2
	// now seed will initialized by same number every time
	// rand.Seed(345)
	// fmt.Println("2. This will depend upon Seed(), print same number every time", rand.Intn(6))
	// 3
	// now it is guranteed that the seed will initialized by random number every time
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("3. This will print random number (1 - 4): ", rand.Intn(6))
	// fmt.Println("3. This will print random number (1 - 5): ", rand.Intn(6)+1)

	// Random form [crpto/rand]
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
