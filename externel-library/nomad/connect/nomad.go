package main

import (
	"fmt"

	"github.com/hashicorp/nomad/api"
)

func main() {
	// Nomad API 클라이언트 설정
	config := api.DefaultConfig()
	config.Address = "http://192.168.20.120:4646"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("Nomad 클라이언트 생성 실패:", err)
		return
	}

	// 리더 정보 조회
	status, err := client.Status().Leader()
	if err != nil {
		fmt.Println("리더 정보 조회 실패:", err)
		return
	}

	fmt.Println("현재 리더 주소:", status)
}
