package main

import (
	"context"
	"fmt"
	"gochat/common"
	chat "gochat/proto"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const PASSWORD = "password"

type Server struct {
	users map[string]string
	lock  sync.RWMutex
}

// main start a gRPC server and waits for connection
func main() {
	s := &Server{
		users: make(map[string]string),
	}
	ctx := context.Background()
	s.Run(ctx)
}

// Run executes the server
func (s *Server) Run(ctx context.Context) error {
	log.SetFlags(0)
	ctx, cancel := context.WithCancel(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1337))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	common.ServerLogf("%s", "Listening on port 1337")
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the chat service to the server
	chat.RegisterChatServer(grpcServer, s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		cancel()
	}

	// broadcast server is running

	<-ctx.Done()
	common.ServerLogf("shutting down")

	grpcServer.GracefulStop()
	return nil
}

func (s *Server) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	if req.Password != PASSWORD {
		return nil, status.Error(codes.Unauthenticated, "wrong password")
	}
	if !s.addUser(req.Name) {
		return nil, status.Error(codes.AlreadyExists, "User already exists")
	}
	common.ServerLogf("%s has logged in", req.Name)

	return &chat.LoginResponse{}, nil
}

func (s *Server) addUser(name string) bool {
	s.lock.Lock()
	if _, ok := s.users[name]; ok {
		return false
	}
	s.users[name] = name
	s.lock.Unlock()
	return true
}
