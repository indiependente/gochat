package main

import (
	"context"
	"fmt"
	common "gochat/common"
	chat "gochat/proto"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(0)
	fmt.Println("Connecting to localhost:1337...")
	conn, err := grpc.Dial("localhost:1337", grpc.WithInsecure())

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unexpected Error: %v", err)
	}

	defer conn.Close()

	c := chat.NewChatClient(conn)
	ctx := context.Background()

	if _, err := c.Login(ctx, &chat.LoginRequest{Name: "pippo", Password: "password"}); err != nil {
		common.Errorf("%v\n", err)
		os.Exit(1)
	}

	common.ClientLogf("%s", "Logged in.")

}
