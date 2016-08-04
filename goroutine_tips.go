package main

// The reason to read the jobs from a channel is that the read from a channel is a blocking operation. 
import (
	"fmt"
	"net"
)

// func (r *RunQueue) runOnce(idx int) {
// 	for {
// 		queueJob, ok := <-r.runQueue
// 		if !ok {
// 			return
// 		}
// 		// run tasks
// 		// [...]
// 	}
// }


// only one goroutine will be running at any one time
func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", message)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	check(err, "srv is ready")

	for {
		conn, err := ln.Accept()
		check(err, "connection accepted")
	}
}
