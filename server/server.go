package main

import (
	"context"
	"fmt"
	"gochat/common"
	chat "gochat/proto"
	"io"
	"time"

	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const password = "password"

type server struct {
	tokens          map[string]string
	users           map[string]*User
	tkLock, usrLock sync.RWMutex
	broadcastCh     chan chat.StreamResponse
}

type User struct {
	name      string
	loginTime time.Time
	stream    chan chat.StreamResponse
}

// main start a gRPC server and waits for connection
func main() {
	s := &server{
		tokens:      make(map[string]string),
		users:       make(map[string]*User),
		broadcastCh: make(chan chat.StreamResponse, 1000),
	}
	ctx := common.SignalContext(context.Background())
	s.Run(ctx)
}

// Run executes the server
func (s *server) Run(ctx context.Context) error {
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

	// start the broadcast goroutine
	go s.broadcast()

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	// broadcast server is running
	<-ctx.Done()

	s.broadcastCh <- chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_ServerShutdown{
			&chat.StreamResponse_Shutdown{},
		},
	}

	close(s.broadcastCh)
	common.ServerLogf("shutting down")

	grpcServer.GracefulStop()
	return nil
}

func (s *server) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	// password check
	if req.Password != password {
		return nil, status.Error(codes.Unauthenticated, "wrong password")
	}
	token := createToken(req.Name)
	u := &User{loginTime: time.Now(), name: req.Name}
	if err := s.addUser(token, u); err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	common.ServerLogf("%s has joined the chat", u.name)
	s.printUsers()
	// return token to user
	return &chat.LoginResponse{Token: token}, nil
}

func (s *server) Logout(ctx context.Context, req *chat.LogoutRequest) (*chat.LogoutResponse, error) {
	fmt.Println("Users")
	s.printUsers()
	u, err := s.getUser(req.Token)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	fmt.Printf("token = %s, user = %v\n", req.Token, u)

	u, err = s.deleteUser(req.Token)
	if err != nil {
		common.Errorf("Delete user with token %s failed: %v\n", req.Token, err)
		return nil, err
	}
	common.ServerLogf("%s has left the chat", u.name)
	return &chat.LogoutResponse{}, nil
}

func (s *server) Stream(stream chat.Chat_StreamServer) error {
	// handle streams
	tk, ok := extractToken(stream.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing token header")
	}
	usr, err := s.getUser(tk)
	if err != nil {
		return status.Error(codes.NotFound, "user not found for token: "+tk)
	}

	// start a new goroutine for this client
	go s.broadcastMessages(stream, tk)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		s.broadcastCh <- chat.StreamResponse{
			Timestamp: ptypes.TimestampNow(),
			Event: &chat.StreamResponse_ClientMessage{&chat.StreamResponse_Message{
				Name:    usr.name,
				Message: req.Message,
			}},
		}
	}

	<-stream.Context().Done()
	return stream.Context().Err()

}
