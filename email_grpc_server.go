package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/iKakalotto/xmail/proto"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmailServer
}

func (s *server) Send(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	golog.Info("Send ... ...")
	return &pb.Response{Success: true}, nil
}

func init() {
	golog.SetTimeFormat(time.DateTime)
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		golog.Fatalf("Read config file failed! Error: %v", err)
	}
}

func StartApplication() {
	port := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		lis, err = net.Listen("tcp", ":0")
		if err != nil {
			golog.Fatalf("Server start failed! Error: %v", err)
		}
	}

	s := grpc.NewServer()
	pb.RegisterEmailServer(s, &server{})
	golog.Infof("Server start on %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		golog.Fatalf("Server start failed! Error: %v", err)
	}
}
