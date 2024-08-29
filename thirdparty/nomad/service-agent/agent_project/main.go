package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("[agent] Start agent...")
	// 실행할 바이너리. 예를 들어 "ls" 명령을 실행한다고 가정
	cmd := exec.Command("/app/service")
	//cmd := exec.Command("./service")
	fmt.Println("[agent] service")

	// 표준 출력을 현재 프로세스의 표준 출력으로 리다이렉트
	cmd.Stdout = os.Stdout

	// 명령어 실행
	if err := cmd.Run(); err != nil {
		fmt.Printf("[agent] 명령어 실행 중 에러 발생: %v\n", err)
		return
	}

	// 프로세스 상태 확인
	if cmd.ProcessState.Success() {
		fmt.Println("[agent] 프로세스가 성공적으로 종료되었습니다.")
	} else {
		fmt.Println("[agent] 프로세스가 에러와 함께 종료되었습니다.")
	}
}
