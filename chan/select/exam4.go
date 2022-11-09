package main

import (
	"fmt"
	"time"
)

func exam4() {
	scheduler := make(chan string, 1)
	prompt := "HAMA"

	select {
	case scheduler <- prompt:
		fmt.Println("이름은: ", <-scheduler)
	case <-time.After(time.Second):
		fmt.Println("시간이 지났습니다.")
	}

	time.Sleep(100 * time.Second)
}
