package main

import (
	"fmt"
	"time"
)

func test_main1() {
	done := make(chan bool)
	defer close(done)

	go func() {
		for i := 1; i < 10; i++ {
			go func() {
				fmt.Println(i)
			}()
		}
		done <- true
	}()

	timer := time.After(time.Duration(1) * time.Second)
	<-timer
	<-done
}

func test_main2() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}
