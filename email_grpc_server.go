package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	pb "github.com/iKakalotto/xmail/proto"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gopkg.in/mail.v2"
)

var (
	gprc_port     int
	mail_port     int
	mail_server   string
	mail_from     string
	mail_account  string
	mail_password string
	env_key_pwd   string
)

type server struct {
	pb.UnimplementedEmailServer
}

func (s *server) Send(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	err := sendMessage(in.Receiver, in.Subject, in.Body)
	if err != nil {
		return &pb.Response{Success: false}, err
	}

	return &pb.Response{Success: true}, nil
}

func init() {
	golog.SetTimeFormat(time.DateTime)
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		golog.Fatalf("Read config file failed! Error: %v", err)
	}

	env_key_pwd = viper.GetString("env.mail.passwd")
	gprc_port = viper.GetInt("server.port")
	mail_port = viper.GetInt("mail.port")
	mail_server = viper.GetString("mail.server")
	mail_account = viper.GetString("mail.account")
	mail_from = viper.GetString("mail.from")
	mail_password = os.Getenv(env_key_pwd)
}

func StartApplication() {
	port := fmt.Sprintf(":%d", gprc_port)
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

func sendMessage(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", mail_from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.AddAlternative("text/html", body)

	d := mail.NewDialer(mail_server, mail_port, mail_account, mail_password)
	return d.DialAndSend(m)
}
