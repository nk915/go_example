package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Prometheus 서버의 엔드포인트 URL
	prometheusURL := "http://localhost:9100/api/v1/query?query=up"

	// HTTP GET 요청 보내기
	response, err := http.Get(prometheusURL)
	if err != nil {
		fmt.Println("HTTP 요청 실패:", err)
		return
	}
	defer response.Body.Close()

	// 응답 데이터 읽기
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("응답 데이터 읽기 실패:", err)
		return
	}

	// JSON 데이터 구조체에 파싱
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("JSON 파싱 실패:", err)
		return
	}

	// 결과 출력
	fmt.Println("응답 데이터:")
	fmt.Println(responseData)
}
