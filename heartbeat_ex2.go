package main

import (
	"time"
	"fmt"
)

func main() {
	cnt := 1
	done := make(chan bool)

	timeout := time.Duration(1)*time.Second
	timer := time.After(timeout)
	go func () {
		for {
			select {
			case <-timer:
				if cnt == 10 {
					done <- true
				} else {
					timer = time.After(timeout)
					fmt.Println("time out", time.Now())
					cnt++
				}
			}
		}
	}()

	<-done
}
