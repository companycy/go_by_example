package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	filename := "/tmp/weather"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	c := "content"
	_, err = io.WriteString(f, c)
	if err != nil {
		fmt.Println(err)
	}

}
