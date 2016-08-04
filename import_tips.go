package main

// import ."package"
// The . (dot) in front of the package name imports the package into your namespace. If you’d import the “fmt” package this way, you could just call Println and other functions without using the fmt. prefix.
import (
	."fmt"
)

func main() {
	Println("hello world")
}
