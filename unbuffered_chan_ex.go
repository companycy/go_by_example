package main

import (
	"fmt"
)

func block1(ch1 chan int) {
	<-ch1
	ch1 <- 1
}

func block2(ch1 chan int) {
	ch1 <- 1
	<-ch1
}

func unblock1(ch1 chan int) {
	go func() {
		ch1 <- 1
	}()

	fmt.Println(<-ch1)
}

func unblock2(ch1 chan int) {
	go func() {
		fmt.Println(<-ch1)
	}()

	ch1 <- 1
}

func main() {
	ch1 := make(chan int)
	// block1(ch1)
	// block2(ch1)

	unblock1(ch1)
	unblock2(ch1)
}
