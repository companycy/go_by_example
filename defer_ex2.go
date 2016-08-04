package main

import "fmt"

func main() {
	fmt.Println("return", a()) // 2
}

func a() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 1
	}()


	return i
}
