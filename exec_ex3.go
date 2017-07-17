package main

import (
	"fmt"
	"log"
	"os/exec"
	"bytes"
)

func main() {
	cmd := exec.Command("python", "-u", "inf_loop.py")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	err := cmd.Wait()
	fmt.Printf("output %s", output.Bytes())
	fmt.Println(err)

	// var out outstream
	// cmd.Stdout = out
	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(cmd.Wait())

}

type outstream struct{}

func (out outstream) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil

}
