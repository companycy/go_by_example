package main

import (
	"fmt"
)

// 6.8 substring, not sub sequence
// x[0...n] and y[0...m]
// 
// T[i, j] = length of longest common substring
// T[i, j] = max {
//               T[i-1, j-1] + 1 if x[i] == y[j],
//               0 otherwise
//           }
//

// a b c f
// b a c a
// => 1

func main() {
	
}
