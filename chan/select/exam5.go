package main

import (
	"fmt"
	"time"
)

var scheduler chan string

func consuming(prompt string) {
	fmt.Println("consuming 호출됨")
	select {
	case scheduler <- prompt:
		fmt.Println("이름을 입력받았습니다 : ", <-scheduler)
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 지났습니다.")
	}
}

func producing(console chan string) {
	var name string
	fmt.Print("이름:")
	fmt.Scanln(&name)
	console <- name
}
func exam5() {
	console := make(chan string, 1)
	scheduler = make(chan string, 1)

	go func() {
		consuming(<-console)
	}()

	go producing(console)

	time.Sleep(100 * time.Second)
}
