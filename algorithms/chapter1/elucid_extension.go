package main

import (
	"fmt"
)

// 1.6
// extension of elucid
func extension_elucide(a, b, N int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}

	x1, y1, d := extension_elucide(a, b, d)
	return y1, (x1-(a/b)*y11), d
}

func main() {
	
}
