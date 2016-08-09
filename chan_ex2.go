// http://golangtutorials.blogspot.com/2011/06/channels-in-go.html

package main

import (
	"fmt"
	"strconv"
	"time"
)

var i int

func makeCake(in chan string) {
	i++
	fmt.Println("Making a cake and sending ...", strconv.Itoa(i))
	in <- strconv.Itoa(i)
}

func makeCake2(in chan string, flavour string) {
	for i := 0; i < 5; i++ {
		// fmt.Println("prepare ", flavour)
		in <- strconv.Itoa(i) + flavour
	}

	close(in)
}

func packCake(out chan string) {
	cake := <-out
	fmt.Println("Packing received cake: ", cake)
}

func packCake2(out ...chan string) {
	for _, i := range out {
		fmt.Println(i)
		go func(j chan string) {
			for l := range j {
				fmt.Println("recv", l)
			}
			// select {
			// case l := <- j: fmt.Println("recv", l)
			// }
		}(i)
	}
}

func packCake3(in1, in2 chan string) {
	open1, open2 := true, true
	var i string
	for {
		if !open1 && !open2 {
			return
		}

		// fmt.Println("wait for new...")
		select {
		case i, open1 = <-in1:
			if !open1 {
				// fmt.Println("0 closed")
			} else {
				fmt.Println(i)
			}
		case i, open2 = <-in2:
			if !open2 {
				// fmt.Println("1 closed")
			} else {
				fmt.Println(i)
			}
		}
	}
}

func test_main() {
	ch := make(chan string)
	for i := 0; i < 3; i++ {
		go makeCake(ch)
		go packCake(ch)
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func main() {
	in1 := make(chan string)
	go makeCake2(in1, "type1")

	in2 := make(chan string)
	go makeCake2(in2, "type2")

	go packCake2(in1, in2)
	// go packCake3(in1, in2)

	time.Sleep(2 * 1e9)
}

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//if both channels are closed then we can stop
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println(" ... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}

func test_main2() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend(choco_cs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend(strbry_cs, "Strawberry", 3) //make 3 strawberry cakes and send

	//one cake receiver and packer
	go receiveCakeAndPack(strbry_cs, choco_cs) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(2 * 1e9)
}
