package main

import (
	"fmt"
)

// 6.6
// ab=b ba=c
// bbbbac
// a = ac
// a = bc
// a = ca
// T(i, j) = T(i, k) * T(k+1, j) for i <= k < j
// => [ T(i, k) = a, T(k+1, j) = c ] = a
// => [ T(i, k) = b, T(k+1, j) = c ] = a
// => [ T(i, k) = c, T(k+1, j) = a ] = a

func main() {
	
}
