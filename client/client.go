package main

import (
	"bufio"
	"context"
	"fmt"
	common "gochat/common"
	chat "gochat/proto"
	"io"
	"log"
	"os"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const tokenHeader = "x-chat-token"

type client struct {
	chat.ChatClient
	name, host, token, password string
}

func main() {
	ctx := common.SignalContext(context.Background())

	log.SetFlags(0)

	name := os.Args[1]

	c := &client{name: name, password: "password", host: "localhost:1337"}
	c.Run(ctx)
}

func (c *client) Run(ctx context.Context) {
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println("Connecting to localhost:1337...")
	conn, err := grpc.DialContext(connCtx, c.host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()

	c.ChatClient = chat.NewChatClient(conn)

	// Login
	loginRes, err := c.Login(ctx, &chat.LoginRequest{Name: c.name, Password: c.password})
	if err != nil {
		common.Errorf("%v\n", err)
		os.Exit(1)
	}

	// save the token
	c.token = loginRes.Token
	common.ClientLogf("Logged in. My token is %s", c.token)

	// start streaming until done
	err = c.stream(ctx)
	if err != nil {
		log.Fatalf(">>> stream failed: %v", err)
	}

	// Logout
	fmt.Println("CTRL+C... Logging you out")
	err = c.logout(ctx, &chat.LogoutRequest{Token: c.token})
	if err != nil {
		common.Errorf("%v\n", err)
		os.Exit(1)
	}
	common.ClientLogf("%s", "Logged out.")
}

func (c *client) stream(ctx context.Context) error {
	md := metadata.New(map[string]string{tokenHeader: c.token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := c.ChatClient.Stream(ctx)
	if err != nil {
		return err
	}
	defer client.CloseSend()

	common.ClientLogf("%s", "connected to stream")

	go c.send(client)
	return c.receive(client)
}

func (c *client) receive(client chat.Chat_StreamClient) error {
	for {
		res, err := client.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			common.Debugf("%s", "stream canceled (usually indicates shutdown)")
			return nil
		} else if err == io.EOF {
			common.Debugf("%s", "stream closed by server")
			return nil
		} else if err != nil {
			return err
		}

		ts := common.TsToTime(res.Timestamp)

		switch evt := res.Event.(type) {
		case *chat.StreamResponse_ClientLogin:
			common.ServerLogf("%s has logged in at %s", evt.ClientLogin.Name, ts)
		case *chat.StreamResponse_ClientLogout:
			common.ServerLogf("%s has logged out at %s", evt.ClientLogout.Name, ts)
		case *chat.StreamResponse_ClientMessage:
			common.MessageLogf("<%s> %s: %s", ts.Format("03:04:05 PM"), evt.ClientMessage.Name, evt.ClientMessage.Message)
		case *chat.StreamResponse_ServerShutdown:
			common.ServerLogf("the server is shutting down")
			syscall.Kill(os.Getpid(), syscall.SIGTERM)
		default:
			common.ClientLogf("unexpected event from the server: %T", evt)
			return nil
		}
	}
}

func (c *client) send(client chat.Chat_StreamClient) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for {
		select {
		case <-client.Context().Done():
			log.Printf("client send loop disconnected")
		default:
			if sc.Scan() {
				err := client.Send(&chat.StreamRequest{Message: sc.Text()})
				if err != nil {
					common.ClientLogf("failed to send message: %v", err)
					return
				}
			} else {
				common.ClientLogf("failed to read input: %s", sc.Err())
				return
			}
		}
	}
}

func (c *client) logout(ctx context.Context, req *chat.LogoutRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.Logout(ctx, req)
	if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
		common.Debugf("unable to logout (connection already closed)")
		return nil
	}

	return err
}
