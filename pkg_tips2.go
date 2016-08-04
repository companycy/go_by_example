// -main.go
// -hello
// --golang
// ---init.go

// main.go
package main

import (
	"fmt"
	"hello/golang"
)

func main() {
	fmt.Println("main func")
	world.Test()
}

// init.go

package world
import (
	"fmt"
)

func init() {
	fmt.Println("init in golang")
}

func localfun() {
	fmt.Println("local func")
}

func Test() {
	localfun()
	fmt.Println("called outside")
}

// _ means only init is called, other func can't be invoked
