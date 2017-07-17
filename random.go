package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(fmt.Sprintf("%v", r.Int31()))
}
