package main

import (
	"fmt"
)

// 1.1
// Multiplication
func multiply(x, y int) int {
	if x == 0 {
		return 0
	}

	z := multiply(x/2, y)
	if (x % 2) == 1 {
		return y + 2*z
	} else {
		return 2 * z
	}
}

func main() {
	fmt.Println(multiply(3, 5))
}
