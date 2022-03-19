package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can start and move 1 spot")
	case 2:
		fmt.Println("You can moce 2 spot")
	case 3:
		fmt.Println("You can moce 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can moce 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can moce 5 spot")
	case 6:
		fmt.Println("You can moce 6 spot and roll dice again")
	default:
		fmt.Println("Whta was that!")
	}
}
