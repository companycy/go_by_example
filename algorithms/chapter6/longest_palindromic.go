package main

import (
	"fmt"
)

// 6.7
// A C G T G T C A A A A T C G
// => A A A A

// T[i, j] = true if s[i...j] is palindromic

// T[i, i] = true since s[i] = s[i]
// T[i, j] = true if s[i] == s[j] and T[i+1, j-1] is true
//           false otherwise

// for i := n; i >= 0; i--
//   for j := n; j >= i; j--


func main() {
	
}
