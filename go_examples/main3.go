package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)

	go func() {
		// Simulating some work here;
		// time.Sleep(2 * time.second)
		// myChannel <- "Hello, World!"
	}()

	// Treat for case when the channel doesn`t respond
	select {
	case msg := <-myChannel:
		fmt.Println("Received:", msg)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out after 10 seconds")
	}
}