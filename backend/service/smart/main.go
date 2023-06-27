package main

import (
	"net"

	"github.com/tinkler/mqttadmin/pkg/logger"
	"github.com/tinkler/ssms/backend/service/smart/internal/gsrv"
	smart_v1 "github.com/tinkler/ssms/backend/service/smart/smart/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger.ConsoleLevel = logger.LL_DEBUG
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	lis, err := net.Listen("tcp", ":10906")
	if err != nil {
		panic(err)
	}
	smart_v1.RegisterSmartGsrvServer(s, gsrv.NewSmartGsrv())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
