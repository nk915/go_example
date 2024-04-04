package main

import (
	"fmt"
	"time"
)

func main() {
	checkTimeout()
	loopTimer()
}

func checkTimeout() {
	// 현재 시간을 가져옵니다.
	currentTime := time.Now()

	// currentTime 로그로 출력
	fmt.Println("Current Time:", currentTime)

	// 지정된 타임아웃 시간(예: 5초) 설정
	timeoutDuration := 5 * time.Second

	// 타임아웃 시간까지 대기
	//<-time.After(timeoutDuration)
	//time.Sleep(timeoutDuration)

	// 타임아웃 시간이 경과했는지 확인
	if time.Since(currentTime) >= timeoutDuration {
		fmt.Println("Timeout occurred!")
	} else {
		fmt.Println("Timeout not occurred.")
	}

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
