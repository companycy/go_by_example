package main

import (
	"fmt"
)

// 6.27
// g -> gap
// S(i, j, g) = Max{ c[0] + c[1]*(g-1)  // S(i-1, j-1, g-1) + diff(i, j)
//                c[0] + c[1](g)  // [i, j-1] or [i-1, j]
// 
// S(i-1, j) => x[0...i-1] and y[0...j], thus s(-, y[j])
// 和答案有些出入, 应该以我为准

func main() {
	
}
