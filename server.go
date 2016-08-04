package main

import (
	"fmt"
	"bufio"
	"net"
	// "time"
)

// only one goroutine will be running at any one time
func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", message)
}


func main() {
	
	ln, err := net.Listen("tcp", ":8080")
	check(err, "srv is  ready")

	for {
		conn, err := ln.Accept()
		check(err, "connect accepted")
		go func() {
			buf := bufio.NewReader(conn)
			for {
				name, err := buf.ReadString('\n')
				if err != nil {
					fmt.Printf("connection broken\n")
					break
				}

				conn.Write([]byte("hello, " + name))
			}
		} ()
	}
}
