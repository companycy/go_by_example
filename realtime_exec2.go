package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

var cmd = exec.Command("/usr/bin/ansible-playbook", "/data/ansible/playbook/authorize_key.yml", "-i", "/data/ansible/inventory/opsbuilderg-0pql1c2h")

func realtimeExec() (int, error) {
	// /usr/bin/ansible-playbook /data/ansible/playbook/authorize_key.yml  -i /data/ansible/inventory/opsbuilderg-0pql1c2h
	var cmd = exec.Command("tailf", "/pitrix/log/ci_server.log")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return 0, err
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		// log.Printf(in.Text()) // write each line to your log, or anything you need
		fmt.Println(in.Text())
	}
	if err := in.Err(); err != nil {
		// log.Printf("error: %s", err)
		fmt.Printf("error: %s", err)
	}

	return 0, nil
}

func nonRealtimeExec() ([]byte, error) {
	// var cmd = exec.Command("/usr/bin/ansible-playbook", "/data/ansible/playbook/test.yml", "-i", "/data/ansible/inventory/opsbuilderg-0pql1c2h")
	var cmd = exec.Command("tailf", "/pitrix/log/ci_server.log")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	var err error
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	err = cmd.Wait()
	fmt.Printf("output: %s", string(output.Bytes()))
	return output.Bytes(), err
}

func main() {
	realtimeExec()
	// nonRealtimeExec()
}
