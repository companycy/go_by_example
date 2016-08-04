package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			fmt.Println("input", num)
			ch <- num
		}
		close(ch)
	}()
	return ch
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			fmt.Println("output", i)
			out <- i * i
		}
		close(out)
	}()
	return out
}

func test_main1() {
	ch1 := gen(2, 3)
	ch2 := sq(ch1)

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(ch))

	for _, i := range ch {
		go func() {
			for n := range i {
				out <- n
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := gen(2, 3)
	out1 := sq(ch1)
	out2 := sq(ch1)

	for i := range merge(out1, out2) {
		fmt.Println("output after merge", i)
	}
}
