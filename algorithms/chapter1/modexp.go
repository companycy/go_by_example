package main

import (
	"fmt"
)

// 1.4
// Modular exponentiation.
func modexp(x, y, N int) int {
	if y == 0 {
		return 1
	}

	z := modexp(x, y/2, N)
	if y % 2 == 1 {
		return x*z*z % N
	} else {
		return z*z % N
	}
}

func main() {
	
}
