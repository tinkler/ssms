package main

import (
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tinkler/mqttadmin/pkg/logger"
	pb_chat_v1 "github.com/tinkler/ssms/backend/service/proxy/chat/v1"
	"github.com/tinkler/ssms/backend/service/proxy/internal/gsrv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	logger.ConsoleLevel = logger.LL_DEBUG

	creds, err := credentials.NewServerTLSFromFile("cert/server_cert.pem", "cert/server_key.pem")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	lis, err := net.Listen("tcp", ":5801")
	if err != nil {
		panic(err)
	}
	pb_chat_v1.RegisterChatGsrvServer(s, gsrv.NewChatGsrv())

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
