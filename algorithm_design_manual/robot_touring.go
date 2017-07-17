package main

import (
	"fmt"
)

func calcDistance(n, nodes) {

}

func robot_touring() {
	for i := 0; i < len(nodes); i++ { // set
		node := nodes[i]
		// node.mark = true // mark it
		minDistance := maxDistance
		var nearestNode Node
		for j := i + 1; j < len(nodes); j++ {
			neighbor := nodes[j]
			if !notSameNode(node, neighbor) {
				if newDistance := calcDistance(node, neighbor); newDistance < minDistance {
					minDistance = newDistance
					nearestNode = neighbor
				}
			}
		}

	}
}

func robot_touring() {
	n := len(P)
	for i := 1; i < n; i++ {
		minDistance = maxDistance
		for distinct_chain(s, t) {
			if newDistance < minDistance {
				new_chain := (sm, tm)
				connect(new_chain)
			}
		}
	}
	connect(P[0], P[n-1])
}

func robot_touring() {
	n = len(P)
	for n := range n! {	// For each of the n! permutations P i of point set P
		getMinDistance(n)
	}
}

func main() {

}
