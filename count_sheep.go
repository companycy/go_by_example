package main

import "time"

func main() {
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		print(i)

	}

}
