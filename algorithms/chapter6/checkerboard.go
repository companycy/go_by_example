package main

import (
	"fmt"
)

// 6.5
// checkerboard: 4 rows and n columns, with integer in each square
// 2n pebbles
// max sum of integer with pebble in square
// 1. all legal patterns = 4 + 3 = 7
// 1/2/3/4
// 13/14/24
// 2.
// M(0) = max {t} for t := range T
// M(1) = max {
//               M(t[k]) + sum(t[i]), t[k] and t[i] are compatible
//               M(t[i-1])
//            }


func main() {
	
}
