package main

import (
	"context"
	"fmt"
	common "gochat/common"
	chat "gochat/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(0)

	name := os.Args[1]

	fmt.Println("Connecting to localhost:1337...")
	conn, err := grpc.Dial("localhost:1337", grpc.WithInsecure())

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unexpected Error: %v", err)
	}

	defer conn.Close()

	c := chat.NewChatClient(conn)
	ctx := context.Background()

	// Login
	loginRes, err := c.Login(ctx, &chat.LoginRequest{Name: name, Password: "password"})
	if err != nil {
		common.Errorf("%v\n", err)
		os.Exit(1)
	}
	common.ClientLogf("Logged in. My token is %s\n", loginRes.Token)

	// do something here
	time.Sleep(5 * time.Second)

	// Logout
	_, err = c.Logout(ctx, &chat.LogoutRequest{Token: loginRes.Token})
	if err != nil {
		common.Errorf("%v\n", err)
		os.Exit(1)
	}
	common.ClientLogf("Logged out.\n")
}
