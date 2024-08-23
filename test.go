package main

import (
	"context"
	"time"

	pb "github.com/iKakalotto/xmail/proto"
	"github.com/kataras/golog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	golog.SetTimeFormat(time.DateTime)
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		golog.Fatalf("Connect gRPC server failed! Error: %v", err)
	}

	defer conn.Close()

	c := pb.NewEmailClient(conn)
	_, _ = c.Send(context.Background(), &pb.Request{})
}
