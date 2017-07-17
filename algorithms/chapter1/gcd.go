package main

import (
	"fmt"
)

// 1.5
// Euclid’s algorithm for finding the greatest common divisor of two numbers.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)	// swap a/b
	// a mod b is done by division, so it's O(n^2) * O(2n) ??
}

func main() {

}
