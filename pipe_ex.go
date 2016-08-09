package main

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, i := range cs {
		go func(in <-chan int) {
			for j := range in {
				output <- j
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func gen(in ...int) <-chan int {
	output := make(chan int)

	go func() {
		for _, i := range in {
			fmt.Println("input", i)
			output <- i
		}
		close(output)
	}()

	return output
}

func sq(in <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for i := range in {
			fmt.Println("output", i*i)
			output <- i * i
		}
		close(output)
	}()

	return output
}

func test_main1() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func merge1(done chan struct{}, cs ...<-chan int) <-chan int {
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, i := range cs {
		go func(in <-chan int) {
			for j := range in {
				select {
				case <-done:
					fmt.Println("recv from done")
				case output <- j:
					fmt.Println("output in merge", j)
				}
			}
			fmt.Println("wg done")
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func test_main2() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	done := make(chan struct{}, 2)
	fmt.Println(<-merge1(done, c1, c2))

	done <- struct{}{}
	done <- struct{}{}
}

func gen1(done <-chan struct{}, nums ...int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output) // VIP
		for _, num := range nums {
			select {
			case output <- num:
			case <-done:
				return
			}
		}
	}()

	return output
}

func sq1(done <-chan struct{}, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output) // VIP
		for i := range input {
			select {
			case output <- i * i:
			case <-done:
				return
			}
		}
	}()

	return output
}

func merge2(done chan struct{}, cs ...<-chan int) <-chan int {
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, i := range cs {
		go func(in <-chan int) {
			defer wg.Done() // VIP
			for j := range in {
				select {
				case <-done:
					return
				case output <- j:
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	done := make(chan struct{})
	defer close(done)

	in := gen1(done, 2, 3)
	c1 := sq1(done, in)
	c2 := sq1(done, in)

	fmt.Println(<-merge2(done, c1, c2))
}
