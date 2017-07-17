package main

import (
	"fmt"
)

func multiplication_divide_conquer(x, y int) {
	// x/y is n-bit int
	// n = max(x's bit, y's bit)
	if n == 1 {
		return x*y
	}

	// Xl = X's left-most n/2, Xr = X's right-most n/2
	// Yl = Y's left-most n/2, Yr = Y's right-most n/2
	P1 = multiplication(Xl, Yl)
	P2 = multiplication(Xr, Yr)
	P3 = multiplication(Xl+Xr, Yl+Yr)
	return 2^n * P1 + P2 + 2^(n/2) * (P3 - P1 - P2)
}

func main() {
	
}
