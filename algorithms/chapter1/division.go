package main

import (
	"fmt"
)

// 1.2
// division
func divide(x, y int) (int, int) {
	if x == 0 {
		return 0, 0
	}

	q, r := divide(x/2, y)
	q, r = 2*q, 2*r
	if x % 2 == 1 {
		r += 1
	}

	if r > y {
		r -= y
		q += 1
	}
	return q, r
}

func main() {
	fmt.Println(divide(11, 3))
}
