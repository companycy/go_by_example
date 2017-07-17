package main

import (
	"fmt"
)

// 6.26
// S(i, j) = Max{ S(i-1, j-1) + s(x[i], y[j])
//                S(i-1, j) + s(-, y[j])
//                S(i, j-1) + s(s[i], -) }
// S(i-1, j) => x[0...i-1] and y[0...j], thus s(-, y[j])
// 和答案有些出入, 应该以我为准

func main() {
	
}
