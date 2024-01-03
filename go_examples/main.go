package main

import (
	"fmt"
	"time"
)

func block1Second(unit int) {
	fmt.Printf("starting unit %d\n", unit)
	end := time.Now().Add(1 * time.Second)
	for time.Now().Before(end) {
	}
	fmt.Printf("finished unit %d\n", unit)
}

func main() {
	for i := 0; i < 5; i++ {
		go block1Second(i)
	}

	time.Sleep(2 * time.Second)
}