package main

import (
	"fmt"
)

// 6.25
// disjoin sets I, J, K
// let S(m, I, J, K, s) be I, J, K, and m is total cnt of elements
// for new a[m+1],
// S(m+1, I, J, K, s) = Or(S(m, I+1, J, K, s), S(m, I, J+1, K, s), S(m, I, J, K+1, s))
// and check
// 1. (I+1, J, K), or (I, J+1, K), or (I, J, K+1) is disjoint
// 2. Sum(I+1) = Sum(J) = Sum(K) or
//    Sum(I) = Sum(J+1) = Sum(K) or
//    Sum(I) = Sum(J) = Sum(K+1)

// S(n, I, J, K, sum)

func main() {
	
}
