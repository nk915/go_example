package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type AgentHash struct {
	ServiceName string    `json:"service_name" validate:"required"`
	TimeStamp   time.Time `json:"time_stamp"`
	Status      string    `json:"status" validate:"required"`
	Hash        string    `json:"hash" validate:"required"`
	Algorithm   string    `json:"algorithm"`
}
type AgentHashResponse struct {
	Result string `json:"result"` // 정보 수신 여부
}

func main() {
	// Create a Resty Client
	client := resty.New()

	req := AgentHash{
		ServiceName: "AA",
		Status:      "OK",
		Hash:        "aa",
	}

	result := AgentHashResponse{}

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&req).
		SetResult(&result). // or SetResult(AuthSuccess{}).
		Post("http://192.168.20.120:5005/hostagent/agenthash")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
