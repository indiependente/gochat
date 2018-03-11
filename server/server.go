package main

import (
	"context"
	"fmt"
	"gochat/common"
	chat "gochat/proto"
	"time"

	"log"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const password = "password"

type Server struct {
	tokens          map[string]string
	users           map[string]*User
	tkLock, usrLock sync.RWMutex
}

type User struct {
	name      string
	loginTime time.Time
}

// main start a gRPC server and waits for connection
func main() {
	s := &Server{
		tokens: make(map[string]string),
		users:  make(map[string]*User),
	}
	ctx := context.Background()
	s.Run(ctx)
}

// Run executes the server
func (s *Server) Run(ctx context.Context) error {
	log.SetFlags(0)

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
	}

	// broadcast server is running

	<-ctx.Done()
	common.ServerLogf("shutting down")

	grpcServer.GracefulStop()
	return nil
}

func (s *Server) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	// password check
	if req.Password != password {
		return nil, status.Error(codes.Unauthenticated, "wrong password")
	}
	token := createToken(req.Name)
	u := &User{loginTime: time.Now(), name: req.Name}
	if err := s.addUser(token, u); err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	common.ServerLogf("%s has joined", req.Name+" ("+token+")")
	// return token to user
	return &chat.LoginResponse{Token: token}, nil
}

func (s *Server) Logout(ctx context.Context, req *chat.LogoutRequest) (*chat.LogoutResponse, error) {
	u, err := s.deleteUser(req.Token)
	if err != nil {
		common.Errorf("%v\n", err)
		return nil, err
	}
	common.ServerLogf("%s has left", u.name)
	return &chat.LogoutResponse{}, nil
}
