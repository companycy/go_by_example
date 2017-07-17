package main

import (
	"fmt"
)

// 6.2
// n hotels, a[1...n]
// v <= 200 mile/day
// x miles => (200-x)^2
// get min cost when stop at a[n]
// p(j) = min cost to stop at a[j]
// p(j) = min{ p(k) + (200-(a[j] - a[k]))^2 } and a[j] - a[k] <= 200 for 0 <= k <= j-1

// => p(j) = min{p(k) +
// k is last position where a[j] - a[k] <= 200

// dag:
// shortest path
// dist(a[j], a[k]) <= 200, then there is an edge

// the minimum penalty for reaching hotel i is found by trying all stopping places for the previous day, adding today's penalty and taking the minimum of those.

func main() {
	
}
