package main

import (
	"fmt"
	"time"
)

func sendValue(ch chan string, value string, id int) {
	fmt.Println("sending value to channel", id)
	time.Sleep((1 * time.Second))
	ch <- value
	fmt.Println("Finished sending value to Channel", id)
}

func receiveValue(ch chan string) {
	val, isChanelOpen := <-ch
	fmt.Println("Recieved value from channel", isChanelOpen, "val:", val)
}

func printCapacityAndLength(ch chan string) {
	// print the capacity of the channel
	fmt.Println("Capacity of channel:", cap(ch))
	// print the length of the channel: // number of values in the channel
	// this will not work for unbuffered channel
	fmt.Println("Length of channel:", len(ch))
}

func isChanelEmpty(ch chan string) {
	// check if the channel is empty or not
	if len(ch) == 0 {
		fmt.Println("Channel is empty")
	} else {
		fmt.Println("Channel is not empty")
	}
}

func isChanelOpen(ch chan string) {
	// check if the channel is open or closed
	// this will also pop the value from the channel
	// if the channel is closed, it will return the value and false
	if _, isOpen := <-ch; isOpen {
		fmt.Println("\n Channel is open")
	} else {
		fmt.Println("\n Channel is closed")
	}
}

func main1() {
	fmt.Println("Welcome to channels in golang")

	myChan := make(chan string, 6) // buffered channel
	// myChan := make(chan string, 2) // buffered channel
	// myChan := make(chan string) // unbuffered channel
	defer close(myChan)

	go sendValue(myChan, "Hello world! - 1", 1)
	go sendValue(myChan, "Hello world! - 2", 2)
	go sendValue(myChan, "Hello world! - 3", 3)
	go sendValue(myChan, "Hello world! - 4", 4)
	go sendValue(myChan, "Hello world! - 5", 5)

	// the order in which the values are sent to the channel is not the order in which they are received
	// because the goroutines are running concurrently
	// and the order of execution is not guaranteed.
	receiveValue(myChan)
	printCapacityAndLength(myChan)
	receiveValue(myChan) // finished sending 4 values in channel and able tore receive 2 values
	printCapacityAndLength(myChan)
	// receiveValue(myChan) // finished sending 5 values in channel and able tore receive 3 values
	// receiveValue(myChan) // finished sending 5 values in channel and able tore receive 4 values
	// receiveValue(myChan) // finished sending 5 values in channel and able tore receive 5 values
	// receiveValue(myChan) // deadlock!

	time.Sleep(3 * time.Second)
	// isChanelOpen(myChan)
	// isChanelEmpty(myChan)
	// printCapacityAndLength(myChan)

	// receive values from channel in a loop
	// this for each loop will continue to read values from channel myChan until and unless the channel is closed
	for val := range myChan {
		fmt.Println("Received value from channel:", val)
		printCapacityAndLength(myChan)

		if len(myChan) == 0 {
			break
		}
	}

	// isChanelOpen(myChan) // panic when channel is empty
	isChanelEmpty(myChan)
	printCapacityAndLength(myChan)

	fmt.Println("Main ends")
}
