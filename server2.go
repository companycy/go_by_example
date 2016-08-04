package main

import (
	"fmt"
	"bufio"
	"net"
	// "time"
)

type ClientJob struct {
	name string
	conn net.Conn
}

// only one goroutine will be running at any one time
func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", message)
}

func genResponse(jobchan chan ClientJob) {
	for {
		job := <-jobchan
		job.conn.Write([]byte("hello, " + job.name))
	}
}

func main() {
	jobchan := make(chan ClientJob)
	go genResponse(jobchan)

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
				jobchan <- ClientJob{name, conn}

				// conn.Write([]byte("hello, " + name))
			}
		} ()
	}
}
