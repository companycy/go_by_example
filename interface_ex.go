package main

import (
	"fmt"
)

type Printf interface {
	CheckPaper(str string)
}

type HPPrint struct {
}

func (p HPPrint) CheckPaper(str string) {
	fmt.Println(str)
}


func test_main() {
	p := HPPrint{}
	p.CheckPaper("test")
}

func main() {
	t := []int {1, 2}
	l := make([]interface{}, len(t))

	for index, val := range t {
		l[index] = val
	}
}
