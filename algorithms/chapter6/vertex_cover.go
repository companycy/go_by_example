package main

import (
	"fmt"
)

// 6.21
// graph G(U, V)
// vertex cover is subset of vertises that includes at least one endpoint of every edge in E.
// input: undirected tree T=(V, E)
// size of smallest vertex cover of T

// from bottom to top:
// S(r) = min{ sum(S(r's children)), 1+sum(S(r's grandchildren)) }

// also needs to record down vertex

func main() {
	
}
