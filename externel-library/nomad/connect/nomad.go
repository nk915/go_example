package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func main() {
	// Nomad API 클라이언트 설정
	config := api.DefaultConfig()
	config.Address = "https://192.168.20.120:4646"
	//config.Address = "https://192.168.20.202:4646"

	// Set up a custom HTTP client with InsecureSkipVerify
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	config.HttpClient = &http.Client{
		Transport: transport,
	}

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("Nomad 클라이언트 생성 실패:", err)
		return
	}

	status(client)
	jobs(client)
}

func status(client *api.Client) {
	// 리더 정보 조회
	status := client.Status()

	leader, err := status.Leader()
	if err != nil {
		fmt.Println("리더 정보 조회 실패:", err)
		return
	}
	fmt.Println("현재 리더 주소:", leader)

	peer, err := status.Peers()
	if err != nil {
		fmt.Println("Peer 정보 조회 실패:", err)
		return
	}
	fmt.Println("현재 Peer 주소:", peer)

}

func jobs(client *api.Client) {
	// Fetch the list of jobs
	jobs, _, err := client.Jobs().List(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Nomad Jobs:")
	for _, job := range jobs {
		fmt.Printf("Job: %s, Status: %s\n", job.Name, job.Status)
		//fmt.Printf("Job: %+v\n", job)
	}
}
