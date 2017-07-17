package main

import (
	"fmt"
)

// 6.1
// input: a[1...n]
// contiguous subsequence of max sum
// 5, 15, -30, 10, -5, 40, 10
// hint: for j := range n, consider contiguous subsequence ends at j
// sub[j] = max{ sub[k, j] } where  0 <= k <= j
// => max { sum[k, j] } where 0 <= k <= j

// solution:
// sub[j] = max{ sub(j-1)+a[j], a[j] }

// implicit dag:
// 5->15->-30 10->-5->40->10


func main() {
	
}
