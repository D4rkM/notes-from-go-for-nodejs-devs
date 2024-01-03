package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello, World!"
	}()

	msg := <- myChannel
	fmt.Println(msg)
}