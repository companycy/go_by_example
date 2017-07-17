package main

import (
	"fmt"
)

// 6.3
// locations m[1...n]
// profits p[1...n]
// m[i] - m[i-1] >= k
// t[j] = max profit from 0 to j
// => max{  }

// t[0, j] = max profits from m[0] to m[j]
// => max{ t[0, k] + t[k+1, j] }  i <= k <= j and m[j] - m[k+1] >= k

// => 

// => max{ t[1] }

// t[0,j] = max profit from m[0] to m[j]
// => max{ t[j-1], t[k]+p[j] } m[j] - m[k] >= k


// if m(i) - m(i-1) >= k => maxProfit(i) = maxProfit(i-1) + p(i)
// else maxProfit(i) = max {
//    maxProfit(i-k) + p(i) where m(i) - m(i-k) >= k
//    p(i) if k doesn't exist
// }

func main() {
	
}
