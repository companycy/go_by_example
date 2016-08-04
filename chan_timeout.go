package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)

	// read from ch
	// go func() {
	// 	ch <- 5
	// }()

	go func() {		// timeout...
		time.Sleep(1000)
		timeout <- true
	}()

	select {
	case <- ch:
		fmt.Println("read from ch")
	case <- timeout:
		fmt.Println("timeout...")
	}
}
