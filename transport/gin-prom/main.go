package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	// Prometheus 레지스트리 생성
	prometheusRegistry := prometheus.NewRegistry()

	// JSON 데이터를 읽어오는 함수
	jsonData, err := readJSONDataFromFile("agent.json")
	if err != nil {
		panic(err)
	}

	// JSON 데이터를 Prometheus 메트릭으로 변환하는 함수
	convertToMetrics(jsonData, prometheusRegistry)

	// Prometheus HTTP 핸들러 등록
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})))

	// 서버 시작
	fmt.Println("Server listening on :8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func readJSONDataFromFile(filename string) (*JSONData, error) {
	// JSON 파일을 읽어서 JSON 데이터 구조체로 변환
	var jsonData JSONData
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}

func convertToMetrics(data *JSONData, registry prometheus.Registerer) {
	// 여기에서 필요한 메트릭 생성 및 레지스트리에 등록
	// 예: CPU, Memory, Disk, Traffic, Process 등의 데이터를 메트릭으로 매핑하여 레지스트리에 등록
	// data.CPU, data.Memory, data.Disk, data.Traffic, data.Process 등의 필드를 사용하여 메트릭 생성
	// 각 필드를 Prometheus 메트릭으로 변환하고 레지스트리에 등록하는 코드를 작성
}
