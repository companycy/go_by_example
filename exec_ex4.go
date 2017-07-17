package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "-u", "inf_loop.py")
	out := outstream{
		outpipe: make(chan string),
	}
	cmd.Stdout = out
	if err := cmd.Start(); err != nil {
		log.Fatal(err)

	}
	fmt.Println("wait", cmd.Wait())

	fmt.Println(<-out.outpipe)
}

type outstream struct {
	outpipe chan string
}

func (out outstream) Write(p []byte) (int, error) {
	s := string(p)
	// fmt.Println(s)
	out.outpipe <- s

	// if s == "" {
		
	// }

	return len(p), nil

}
