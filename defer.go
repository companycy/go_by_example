package main

import ."fmt"

func p(i int) {
	Println(i)
}

func main() {
	a := []int{1, 2, 3}
	for _, i := range a {
		Println(i)
		defer p(i)	// first in, last out
	}
}
