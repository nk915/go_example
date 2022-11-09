package main

import (
	"fmt"
	"time"
)

func exam1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			c1 <- "one"
		}
	}()
	go func() {
		for {
			time.Sleep(10 * time.Second)
			c2 <- "two"
		}
	}()

	for {
		fmt.Printf("\n --> start select \n")
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
		fmt.Printf(" --> end select \n\n")
	}
}
