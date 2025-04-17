package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("John")
	println("User ID:", id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)
	wg.Add(2)
	// defer close(ch)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	// wg.Done()
	close(ch)

	for msg := range ch {
		log.Println("Message:", msg)
	}

	log.Println(time.Since(now))
}

func getUserByName(name string) string {
	time.Sleep(1 * time.Second)

	return fmt.Sprintf("%s-2", name)
}

// func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
func getUserChats(id string, ch chan *Message, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	ch <- &Message{chats: []string{"Chat1", "Chat2"}}
	wg.Done()
}

// func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
func getUserFriends(id string, ch chan *Message, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	ch <- &Message{friends: []string{"Friend1", "Friend2"}}
	wg.Done()
}

func leaky() {
	// in this example, no one is writing to the channel, so it will block
	// and the program will not terminate
	ch := make(chan int)
	// defer close(ch) // solution

	go func() {
		msg := <-ch
		fmt.Println("Received:", msg)
	}()
}
