package main

import (
	"fmt"
	"time"
)

func getMessage(messageChannel chan<- string) {
	time.Sleep(2 * time.Second)
	messageChannel <- "My message"
}

func main() {
	messageChannel := make(chan string)
	go getMessage(messageChannel)

	select {
	case msg := <-messageChannel:
		fmt.Println(msg)
	default:
		fmt.Println("no message received")
	}
}
