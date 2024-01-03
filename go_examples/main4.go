package main

import (
	"fmt"
)

func main() {
	var ch1 chan string = make(chan string)
	var ch2 chan int = make(chan int)

	go func() { for { ch1 <- "Str" } }()
	go func() { for i := 0; ; i++ { ch2 <- i } }()
	// Smart behaviour, the select will act randomly 
	// in case of concurrent channels, preventing starvation
	var cnt int = 0
	for cnt < 5 {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
			cnt++
		case m2 := <-ch2:
			fmt.Println(m2)
			cnt++
		}
	}
	fmt.Println("Done.")
}