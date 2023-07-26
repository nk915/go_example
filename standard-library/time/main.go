package main

import (
	"fmt"
	"time"
)

func main() {
	loopTimer()
}

func loopTimer() {
	fmt.Println("--> start function..")
	ticker := time.NewTicker(time.Second)
	timer := time.After(12 * time.Second)

	loop := true
	for loop {

		select {
		case <-ticker.C:
			fmt.Println(" .")
		case <-timer:
			fmt.Println("12 seconds elapsed!")
			return
		}
		fmt.Println("loop")
	}

	fmt.Println("--> end function..")
}
