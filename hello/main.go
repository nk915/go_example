package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("--> run ")

	// 호스트 이름을 얻기
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("hostname err: %v\n", err)
		return
	}

	fmt.Printf("args: %v\n", os.Args)
	fmt.Printf("hostname: %s\n", hostname)

	for {
		fmt.Println("loop")
		time.Sleep(10 * time.Second)
	}
}
