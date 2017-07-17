package main

import (
	"fmt"
)

// 6.22
// a[1...n]
// sum = t
// subproblem: sum(a[1...i]) = s

// S(s, i+1) = f(S(s, i))
// => Or(S(s-a[i+1], i), S(s, i))

// for 0 -> t
//   for 1 -> n
//
// S(t, n)
// -> O(t*n)


func main() {
	
}
