package main

import (
	"flag"
	"fmt"
	pb "grpc/api"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 2001, "The server port")
)

type clientSession struct {
	stream pb.ChatService_ConnectServer
	id     string
	active bool
	err    chan error
}

type server struct {
	pb.UnimplementedChatServiceServer // GRPC 기본 서버 인터페이스로 임베드 서버 구현할 때 사용
	clients                           map[string]*clientSession
	mu                                sync.Mutex
}

func newServer() *server {
	return &server{
		clients: make(map[string]*clientSession),
	}
}

func (s *server) Connect(stream pb.ChatService_ConnectServer) error {
	errChan := make(chan error)
	clientId := ""

	// 클라이언트 스트림 종료 감지 및 에러처리를 위한 고루틴 시작
	go s.monitorClient(stream, errChan, &clientId)

	for {
		req, err := stream.Recv()
		if err != nil {
			errChan <- err
			return err
		}

		if clientId == "" {
			clientId = strings.TrimSpace(req.ClientId) // 공백 및 개행 문자 제거
			s.addClient(clientId, stream, errChan)
			s.sendMessageToClientById(clientId, "Welcome, "+clientId)
		} else {
			log.Printf("Received message from %s : %s", clientId, req.Message)
		}
	}
}

func (s *server) monitorClient(stream pb.ChatService_ConnectServer, errChan chan error, clientId *string) {
	defer s.removeClient(*clientId)
	for {
		select {
		case <-stream.Context().Done():
			log.Printf("Client disconnected: %s\n", *clientId)
			return
		case err := <-errChan:
			if err != nil {
				s.logError(*clientId, err)
			}
			return
		}
	}
}

func (s *server) addClient(clientID string, stream pb.ChatService_ConnectServer, errChan chan error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clients[clientID] = &clientSession{
		stream: stream,
		id:     clientID,
		active: true,
		err:    errChan,
	}
	log.Printf("Client connected: %s\n", clientID)
}

func (s *server) removeClient(clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, clientID)
}

func (s *server) logError(clientID string, err error) {
	if clientID != "" {
		log.Printf("Error with client %s: %v\n", clientID, err)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func (s *server) sendMessageToClient(client *clientSession, message string) {
	if client != nil && client.active {
		log.Printf("send to message : %s", message)
		if err := client.stream.Send(&pb.ServerMessage{Message: message}); err != nil {
			client.err <- err
		}
	} else {
		log.Printf("send fail %s", message)
	}
}

func (s *server) sendMessageToClientById(clientId, message string) {
	s.mu.Lock()
	client := s.clients[clientId]
	s.mu.Unlock()

	s.sendMessageToClient(client, message)
}

func (s *server) broadcastMessage(message string) {
	for idx := range s.clients {
		s.sendMessageToClient(s.clients[idx], message)
	}
}

func (s *server) allClose() {

}

func main() {
	flag.Parse()

	// GRPC 서버 리스너 생성
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	// GRPC 서버 생성
	grpcServer := grpc.NewServer()
	grpcServerService := newServer()

	// Stream 서비스 등록
	pb.RegisterChatServiceServer(grpcServer, grpcServerService)

	// 10초 후 클라이언트에 메시지를 브로드캐스트
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Printf("broadcast message")
			grpcServerService.broadcastMessage("This is a server-initiated message to all clients.")
		}
	}()

	// 서버 시작
	log.Printf("server listening at %v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

}
