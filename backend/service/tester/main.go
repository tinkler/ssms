package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	mrz_v1 "github.com/tinkler/mqttadmin/mrz/v1"
	"github.com/tinkler/mqttadmin/pkg/logger"
	"github.com/tinkler/ssms/backend/service/tester/internal/gsrv"
	pb_tls_v1 "github.com/tinkler/ssms/backend/service/tester/tls/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	logger.ConsoleLevel = logger.LL_DEBUG

	cert, err := tls.LoadX509KeyPair("cert/server_cert.pem", "cert/server_key.pem")
	if err != nil {
		panic(err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
	})

	// creds, err := credentials.NewServerTLSFromFile("cert/server_cert.pem", "cert/server_key.pem")
	// if err != nil {
	// 	log.Fatalf("failed to create credentials: %v", err)
	// }

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb_tls_v1.RegisterTlsGsrvServer(grpcServer, gsrv.NewTlsGsrv())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	// client
	caCert, err := ioutil.ReadFile("cert/ca_cert.pem")
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("credentials: failed to append certificates")
	}
	clientCreds := credentials.NewTLS(&tls.Config{
		RootCAs: caCertPool,
	})
	// clientCreds, err := credentials.NewClientTLSFromFile("cert/ca_cert.pem", "x.test.example.com")
	// if err != nil {
	// 	log.Fatalf("failed to load credentials: %v", err)
	// }

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		log.Fatalf("did no connect: %v", err)
	}
	defer conn.Close()

	rgc := pb_tls_v1.NewTlsGsrvClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	chatStream, err := rgc.ChatInRoom(ctx)
	if err != nil {
		log.Fatalf("can not make a in room chatting, %v", err)
	}

	data := new(pb_tls_v1.Chat)
	go func() {
		for {
			any, err := chatStream.Recv()
			if err == io.EOF || status.Code(err) == codes.Canceled {
				break
			}
			if err != nil {
				logger.Error(err)
			}
			res := mrz_v1.ToDataTypedRes[*pb_tls_v1.Chat](any)
			logger.Info("%v", res.Data.Messages)
			data = res.Data
		}
	}()

	for {
		req := mrz_v1.NewDataTypedModel[*pb_tls_v1.Chat]()
		req.Data = data
		err := chatStream.Send(req.ToAny())
		if err == io.EOF || status.Code(err) == codes.Canceled {
			return
		}
		if err != nil {
			logger.Error(err)
		}
	}
}
