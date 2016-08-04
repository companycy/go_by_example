package main

import (
	"fmt"
	"time"
)

func main() {
	durationToBlowUp := time.Duration(5) * time.Second

	timer := time.NewTimer(durationToBlowUp)

	go func() {
		<-timer.C
		fmt.Println("Timer blows!")
	}()

	// instead of 5 seconds, make it faster
	// reset bomb to 2 seconds!!!!

	reset := timer.Reset(time.Duration(2) * time.Second)

	// ok go to sleep and wait for the pop sound
	time.Sleep(time.Duration(6) * time.Second)

	fmt.Println("Manage to reset timer from 5 to 2 seconds? : ", reset)

}
