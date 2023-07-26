package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleWebSocket(c *gin.Context) {
	// WebSocket 연결 업그레이드
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket 연결 실패:", err)
		return
	}
	defer conn.Close()

	for {
		// 클라이언트로부터 메시지 읽기
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("메시지 수신 오류:", err)
			break
		}

		// 받은 메시지를 출력
		fmt.Printf("받은 메시지: %s\n", message)

		// 클라이언트에게 동일한 메시지 보내기
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("메시지 전송 오류:", err)
			break
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/ws", handleWebSocket)
	fmt.Println("웹소켓 서버 시작. ws://localhost:8080/ws 로 접속하세요.")
	r.Run(":8080")
}
