package main

import (
	"fmt"
	// "time"
)

func producer() {
	fmt.Println("start producer goroutine")
}

func consumer() {
	fmt.Println("start consumer goroutine")
}

func main() {
	c := make(chan int)
	defer close(c)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			c <- i
			// time.Sleep(time.Second * 1)
		}
		done <- true
	}()

	go func() {
		for {
			<-c
			fmt.Println("consumer reads out")
		}
	}()

	<-done
}
