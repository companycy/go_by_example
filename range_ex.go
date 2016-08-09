package main

import (
	"fmt"
	"time"
)

func input(in chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("prepare", i)
		in <- i
	}
	close(in)
}

func output(out chan int) {
	for i := range out {
		fmt.Println("out", i)
	}
}

func main() {
	in := make(chan int)
	go input(in)
	go output(in)

	time.Sleep(2 * 1e9)
}
