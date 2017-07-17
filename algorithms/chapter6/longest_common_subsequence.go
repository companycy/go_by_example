package main

import (
	"fmt"
)

// 6.11 subsequence, not substring
// L[i, j] = if x[i] == y[j], L[i, j] = L[i-1, j-1] + 1
//           otherwise max{ L[i-1, j], L[i, j-1] }

func main() {
	
}
