package main

import (
	"fmt"
)

// 6.23
// P(i, b) = max reliability for M[0...i] with budget b
// M[i] cost is c[i], and reliability is r[i]
// assume t = b/c[i]
// P(i+1, b) = max{ P(i, b)*r[i], P(i, b-k*c[i])*(1-(1-r[i])^k) } where k := range t
// P(n, B)

// for b := range B
//   for i := range n
//     t = b/c[i]
//     for k := range t
//       P(i, b) = max{} // ...

func main() {
	
}
