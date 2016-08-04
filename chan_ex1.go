package main

import (
	"fmt"
)

func test_main() {
	ch1 := make(chan int)
	defer close(ch1)

	done := make(chan bool)
	defer close(done)

	go func() {
		for i := range []int{1, 2} {
			fmt.Println("input", i)
			ch1 <- i
		}
	}()

	go func() {
		for i := 0; i < 2; i++ {
			select {
			case i := <-ch1:
				fmt.Println("output", i)
			}
		}
		done <- true
	}()

	<-done
}

func main() {
	ch := make(chan int)
	close(ch)

	fmt.Println(<-ch)
}
