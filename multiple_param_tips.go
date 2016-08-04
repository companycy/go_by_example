
package main

import (
	"fmt"
)

func MyPrint(str ...string) {
	for _, s := range str {
		fmt.Println(s)
	}
}

func main() {
	MyPrint("hello", "golang")
}
