package main

import (
	pb "assignment3/chittychat"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type chittychatserver struct {
	pb.UnimplementedChittyChatServer
	mu          sync.Mutex
	lamportTime uint64
	clients     map[string]pb.ChittyChat_JoinChatServer
}

func (s *chittychatserver) JoinChat(req *pb.JoinRequest, stream pb.ChittyChat_JoinChatServer) error {
	s.mu.Lock()
	s.clients[req.Username] = stream
	s.lamportTime++
	currentTime := s.lamportTime
	log.Printf("%v enter the chat system", req.Username)
	s.mu.Unlock()

	s.broadcast(&pb.ChatMessage{
		Username:  "server",
		Content:   "Participant " + req.Username + " joined Chitty-Chat at Lamport time " + strconv.FormatUint(currentTime, 10),
		Timestamp: currentTime,
	})

	log.Printf("after broadcast %d enter the chat system", s.lamportTime)
	<-stream.Context().Done()
	return nil

}

func (s *chittychatserver) LeaveChat(ctx context.Context, req *pb.LeaveRequest) (*pb.Empty, error) {
	s.mu.Lock()
	if _, exists := s.clients[req.Username]; !exists {
		s.mu.Unlock()
		return nil, status.Errorf(codes.NotFound, "username not found")
	}
	delete(s.clients, req.Username)
	s.lamportTime++
	currentTime := s.lamportTime
	log.Printf("%v leave the chat system", req.Username)
	s.mu.Unlock()

	s.broadcast(&pb.ChatMessage{
		Username:  "System",
		Content:   "Participant " + req.Username + " left Chitty-Chat at Lamport time " + strconv.FormatUint(currentTime, 10),
		Timestamp: currentTime,
	})

	log.Printf("after broadcast %d enter the chat system", s.lamportTime)
	return &pb.Empty{}, nil
}

func (s *chittychatserver) broadcast(msg *pb.ChatMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for username, stream := range s.clients {
		if err := stream.Send(msg); err != nil {
			log.Printf("Failed to send message to %s: %v", username, err)
			delete(s.clients, username)
		}
	}
	log.Printf("The message has been broadcasted")
}

func (s *chittychatserver) SendMessage(ctx context.Context, msg *pb.ChatMessage) (*pb.Empty, error) {
	s.mu.Lock()
	s.lamportTime = max(msg.Timestamp, s.lamportTime) + 1
	msg.Timestamp = s.lamportTime
	log.Printf("%v sent the message", msg.Username)
	s.mu.Unlock()

	s.broadcast(msg)
	log.Printf("after broadcast %d enter the chat system", s.lamportTime)
	return &pb.Empty{}, nil
}

func main() {
	logFile, err := os.OpenFile("../log/chat_server.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile) // 将日志输出设置为文件

	log.Println("Starting chat server...") // Log startup

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../chitty-chat-tls/server.crt", "../chitty-chat-tls/server.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterChittyChatServer(s, &chittychatserver{
		clients: make(map[string]pb.ChittyChat_JoinChatServer),
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
