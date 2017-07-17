package main

import (
	"fmt"
)

// 6.4
// s[1...n]
// ex. itwasthebestoftimes
// dict(w) can be used to determine string s is a word or not
// determine whether s[1...n] can be reconstituted to a sequence of valid words

// T[i] = Or{ T[i-k] } for all dict(s[i-k...i]) is true where 0 <= k < i

func main() {
	
}
