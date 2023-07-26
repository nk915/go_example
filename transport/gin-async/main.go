package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/resource", func(c *gin.Context) {
		// Client로부터 요청에 대한 처리

		// 비동기 처리를 위해 고루틴 생성
		go func() {
			// 별도의 비동기 작업 수행
			time.Sleep(5 * time.Second)

			// 비동기 작업이 완료된 후 응답을 보내기 위해 원래의 요청과 같은 Context를 사용
			fmt.Println("비동기 처리가 완료되었습니다.")
			c.JSON(200, gin.H{
				"message": "비동기 처리가 완료되었습니다.",
			})
		}()

		// 원래의 요청에 대한 응답 보내기
		fmt.Println("요청에 대한 응답입니다.")
		c.JSON(200, gin.H{
			"message": "요청에 대한 응답입니다.",
		})
	})

	r.Run(":8080")
}
