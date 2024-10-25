package main

import (
	pb "assignment3/chittychat"
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	username    = flag.String("username", "anonymous", "Your chat username")
	serverAddr  = flag.String("server", "localhost:50051", "The server address in the format of host:port")
	mu          sync.Mutex
	lamportTime uint64
)

func main() {
	flag.Parse()
	// 设置日志文件
	logFilePath := fmt.Sprintf("../log/%s.log", *username)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile) // 将日志输出设置为文件

	log.Println("Starting chat client...") // Log startup
	creds, err := credentials.NewClientTLSFromFile("../chitty-chat-tls/server.crt", "")
	if err != nil {
		log.Fatalf("Failed to load TLS certificate: %v", err)
	}
	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)
	stream, err := client.JoinChat(context.Background(), &pb.JoinRequest{Username: *username})
	if err != nil {
		log.Fatalf("Could not join chat: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error receiving message: %v", err)
				return
			}
			mu.Lock()
			lamportTime = max(lamportTime, msg.Timestamp) + 1
			mu.Unlock()
			log.Printf("[%d] %s: %s\n", lamportTime, msg.Username, msg.Content)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "/quit" {
			break
		}
		mu.Lock()
		lamportTime++
		mu.Unlock()
		_, err := client.SendMessage(context.Background(), &pb.ChatMessage{
			Username:  *username,
			Content:   message,
			Timestamp: lamportTime,
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}

	_, err = client.LeaveChat(context.Background(), &pb.LeaveRequest{Username: *username})
	if err != nil {
		log.Printf("Error leaving chat: %v", err)
	}
}
