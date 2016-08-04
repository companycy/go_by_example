package main

import (
	"fmt"
)

func SayTwice(text string) string {
	return text + text
}

func main() {
	fmt.Println(SayTwice("Hello"))

	header := []byte{1, 2, 3, 4}
	fmt.Println(header)
	fmt.Println(len(header))

	header = make([]byte, 4)
	fmt.Println(header)
}
