package main

import (
	"fmt"
	"time"
)

func exam2() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1 * time.Second)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}

		scheduling()
	}
}

func process(ch chan string) {
	time.Sleep(10 * time.Second)
	ch <- "process successful"
}

func scheduling() {
	//do something
	fmt.Println("scheduling")
}
