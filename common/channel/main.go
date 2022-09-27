package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {

	fmt.Printf("run square\n")
	loop := 0
	for {
		fmt.Printf("\n\n\n loop %d\n", loop)
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n)
		case <-quit:
			wg.Done()
			return
		}
		loop++
	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		fmt.Println("insert data ", i)
		ch <- i
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 10)
	quit <- true
	wg.Wait()
}
