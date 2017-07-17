package main

import (
	"fmt"
)

// 6.10
// T(i, j) probability for j head in total i coins are tossed

// T(1, 1) = p1
// T(1, 0) = 1-p1

// T(2, j) = T(1, j-1) * p2 + T(1, j) * (1-p2)
// T(3, j) = T(2, j-1) * p3 + T(1, j) * (1-p3)

// T(i, j) = T(i-1, j-1) * pj + T(i-1, j) * (1-pj)


func main() {
	
}
