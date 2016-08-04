package main

// func AuthenticateRequest(r *Request) error {
// 	err := authenticate(r.User)
// 	if err != nil {
// 		return fmt.Errorf("authenticate failed: %v", err)
// 	}
// 	return nil
// }

import "fmt"

type Test interface {
	Error() string
}

func get_nil() error {
	var e Test = nil
	return e
}

func main() {
	if get_nil() == nil {
		fmt.Println("get nil")
	} else {
		fmt.Println("not real nil")
	}
}
