package main

import (
	"fmt"
)

// 6.24
// O(m*n) space and time
// -> O(n) space

// P(i, j) = max{ P(i-1, j-1) if a[i] == b[j],
//                1+P(i-1, j-1) else
//                1+P(i-1, j),
//                1+P(i, j-1)} 
//          

// only two lines are required: P(i-1) and P(i)
// thus, O(n) space is required instead of O(m*n)

// 

func main() {
	
}
