package main

import (
	"fmt"
	"strconv"
	"time"
)

type Job struct {
	j_id   int
	j_name string
}

func test_main() {
	cnt := 3
	ch := make(chan bool, cnt)
	defer close(ch)

	for {
		ch <- true
		go func() {
			fmt.Println("one more goroutine")
			time.Sleep(time.Duration(5) * time.Second)
			// do sth
			<-ch
		}()
	}

}

func main() {
	// cnt := 3
	// done := make(chan bool, cnt)
	// defer close(done)

	job_pool := make(chan Job) // unbuffered
	go func() {
		for i := 0; i < 5; i++ { // more than pool size
			job := Job{
				i, strconv.Itoa(i),
			}
			job_pool <- job
			fmt.Println("input", job)
		}
	}()

	go func() {
		for j := range job_pool {
			fmt.Println("output", j)
			// do sth
			// time.Sleep(time.Duration(5) * time.Second)
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
	// <-done
}
