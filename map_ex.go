package main

import (
	"fmt"
)

func test(m map[string]string) {
	m["a"] = "b"
}

func main() {
	m := make(map[string]string)
	fmt.Println(m)
	test(m)
	fmt.Println(m)
}
