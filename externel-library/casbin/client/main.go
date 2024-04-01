package main

import (
	"context"
	"encoding/base64"
	"log"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := auth.NewAuthorizationClient(conn)

	headerValue := "Bearer " + getJwt()
	request := &auth.CheckRequest{
		Attributes: &auth.AttributeContext{
			Request: &auth.AttributeContext_Request{
				Http: &auth.AttributeContext_HttpRequest{
					Method:  "GET",
					Path:    "/alice_data/1",
					Headers: map[string]string{"authorization": headerValue},
				},
			},
		},
	}

	response, err := c.Check(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling Check API: %s", err)
	}
	log.Printf("Response from server: %v", response)
}

func getJwt() string {
	Jwt := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2OTcwNzcyMDcsImV4cCI6MTY5NzA3ODQwOSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoiYWxpY2UiLCJHaXZlbk5hbWUiOiJKb2hubnkiLCJTdXJuYW1lIjoiUm9ja2V0IiwiRW1haWwiOiJqcm9ja2V0QGV4YW1wbGUuY29tIiwiUm9sZSI6WyJNYW5hZ2VyIiwiUHJvamVjdCBBZG1pbmlzdHJhdG9yIl19.KkD8kvbAEc3eMIZLnNr8BRQ7n7weFyShdQrpu7kr8RI"

	return base64.StdEncoding.EncodeToString([]byte(Jwt))
}
