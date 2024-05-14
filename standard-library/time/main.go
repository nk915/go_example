package main

import (
	"fmt"
	"time"
)

func main() {
	compareTime()
	//zeroTime()
	//checkTimeout()
	//loopTimer()
}

func nowMillisecond() string {
	// 현재 시간을 millisecond 단위로 얻기
	now := time.Now().UnixNano() / int64(time.Millisecond)

	// Unix 밀리초로부터 시간 생성
	t := time.Unix(0, now*int64(time.Millisecond))

	// 시간을 문자열로 변환
	return t.Format("2006-01-02 15:04:05.999")
}

func compareTime() {
	//bTimeString := "2024-04-19 10:30:58.755"
	//bTimeString := "2024-04-19 10:30:58"

	// KST 시간대 생성
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("시간대 로드 에러:", err)
		return
	}

	//bTimeString := time.Now().In(loc).Format("2006-01-02 15:04:05")
	bTimeString := time.Now().In(loc).Format("2006-01-02 15:04:05.999")

	//bTime, err := time.ParseInLocation(time.DateTime, bTimeString, loc)
	bTime, err := time.ParseInLocation(time.DateTime, bTimeString, loc)
	if err != nil {
		fmt.Println("두 번째 날짜 문자열 파싱 에러:", err)
		return
	}

	fmt.Printf("str:%s, parse_str:%s\n", bTimeString, bTime.String())

	cur := time.Now().In(loc)
	cDur := bTime.Sub(cur)
	fmt.Println(cur.String(), bTime.String())

	fmt.Println("b -> a 대입 시간차: ", cDur)

}

type MyStruct struct {
	SomeTime time.Time
}

// SetZero 함수는 SomeTime 필드를 제로값으로 설정합니다.
func (a *MyStruct) SetZero() {
	a.SomeTime = time.Time{}
}

// IsZero 함수는 SomeTime 필드가 제로값인지 확인합니다.
func (a *MyStruct) IsZero() bool {
	return a.SomeTime.IsZero()
}

func zeroTime() {
	var myStruct MyStruct

	// SomeTime을 현재 시간으로 설정
	myStruct.SomeTime = time.Now()
	fmt.Println("Before reset:", myStruct.SomeTime)

	// SetZero 함수를 호출하여 SomeTime을 제로값으로 설정
	myStruct.SetZero()
	fmt.Println("After reset:", myStruct.SomeTime)

	// IsZero 함수를 호출하여 SomeTime이 제로값인지 확인
	if myStruct.IsZero() {
		fmt.Println("SomeTime is zero value")
	} else {
		fmt.Println("SomeTime is not zero value")
	}
}

func checkTimeout() {
	// 현재 시간을 가져옵니다.
	currentTime := time.Now()

	// currentTime 로그로 출력
	// fmt.Printf("Current Time: %s\n", currentTime.GoString())
	fmt.Printf("Current Time: %s\n", currentTime.String()) // same: fmt.Println("Current Time:", currentTime)

	// 지정된 타임아웃 시간(예: 5초) 설정
	timeoutDuration := 5 * time.Second
	fmt.Printf("timeout:%s\n", timeoutDuration.String())

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
