package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("○ Welcome to channels in golang")
	// intro()

	// ch is channel, capable of sending and receiving value of type int
	// ch := make(chan int)

	// [1] .... [6]
	// send and receive type
	// go receiver(ch)

	//[1]
	// writing value in ch variable
	// ch <- 42

	//[2]
	// ch <- 42
	// close(ch)
	// ch <- 42

	// [3]
	// ch <- 45
	// close(ch)

	// [4] & [5]
	// ch <- 45
	// ch <- 34
	// ch <- 56
	// ch <- 50
	// ch <- 49
	// ch <- 109
	// close(ch)

	// [6] receiving and sending value through channel
	// ch <- 42
	// fmt.Println(<-ch)
	// close(ch)

	// [7] ---------------------------------------------------------
	// receive-only type
	// go receiver2(ch)

	// [7] Only sending value through channel
	// ch <- 12342
	// close(ch)

	// [8] ---------------------------------------------------------
	// send-only type
	// go sender(ch)

	// [8] Only receiving value through channel
	// fmt.Println(<-ch)
	// close(ch)

	// [9]
	// ch_4 := make(chan int) // deadlock it can only recieve 1 value
	// ch_4 := make(type, size of the channel to hold number of values)
	ch_4 := make(chan int, 2) // this chan can hold 2 values at any interval of time

	go receiver_4(ch_4)
	ch_4 <- 1 // this val will be consumed by buffer "receiver_4: fmt.Println(<-ch)"
	ch_4 <- 2
	ch_4 <- 3
	// ch_4 <- 4 // overflow the capacity of the buffer
	// fmt.Println(`This shows us that a goroutine will blocked when the buffer of a buffer channel is full`)

	time.Sleep(100 * time.Millisecond)
}

// [9]
func receiver_4(ch <-chan int) {
	fmt.Println(<-ch)
}

func receiver(ch chan int) {
	// [1]
	// fmt.Println(<-ch)

	// [2]
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	//[3] & [4]
	// receiving value from the channel infinitily
	// for {
	// 	i, ok := <-ch
	// 	if ok == false {
	// 		fmt.Println(i, ok, "<-- loop broken")
	// 		break //exit break loop
	// 	} else {
	// 		fmt.Println(i, ok)
	// 	}
	// }

	// [5]  doing the above same task with for each
	// this for each loop will continue to read values from channel ch until and unless the channel is closed
	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// [6] receiving and sending value through channel
	fmt.Println(<-ch) // receiving
	ch <- 12342       // sending

	fmt.Println("Channel has been closed by sender")
}

//[7]
// unidirectional channel
// Receiver2 is only able to receive values not able to send values
func receiver2(ch <-chan int) {
	fmt.Println(<-ch) //receiving
	// ch <- 12342 //err: invalid operation: ch <- 12342 (send to receive-only type <-chan int)
	fmt.Println("Receiver2: Channel has been closed by sender")
}

//[8]
// unidirectional channel
// Sender is only able to send values not able to receive values
func sender(ch chan<- int) {
	// fmt.Println(<-ch) //err: invalid operation: <-ch (receive from send-only type chan<- int)
	ch <- 99 //sending
	fmt.Println("Sender: Channel has been closed by sender")
}

func intro() {
	fmt.Println(`
	○ Channel is a communication object using which go routines with each other 
	○ Channel is also considered as a pipe which transfers data between go routines
	○ Go routines can write data into this channel or this pipe and other go routines 
	  could read data from the other end of this pipe.
	○ we use inbuilt make() function to make a channel
	○ Channels are static typed
	
	// syntax: reding the latest value of ch
	// <-ch
	// fmt.Println(<-ch)


	// syntax: closing a channel
	// close(ch)

	`)
}
