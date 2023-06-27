package gsrv

import (
	"context"
	"net"
	"testing"

	mrz_v1 "github.com/tinkler/mqttadmin/mrz/v1"
	pb_chat_v1 "github.com/tinkler/ssms/backend/service/proxy/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

func createServer() (func(), error) {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return nil, err
	}
	pb_chat_v1.RegisterChatGsrvServer(s, NewChatGsrv())
	go s.Serve(lis)
	return func() { s.Stop() }, nil
}

func createClient() (pb_chat_v1.ChatGsrvClient, func(), error) {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	c := pb_chat_v1.NewChatGsrvClient(conn)
	return c, func() { conn.Close() }, nil
}

func TestChat(t *testing.T) {

	closeServerFunc, err := createServer()
	if err != nil {
		t.Fatal(err)
	}
	defer closeServerFunc()
	c, closeClientFunc, err := createClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closeClientFunc()

	req := mrz_v1.NewTypedModel[*pb_chat_v1.Chat, *structpb.Struct]()
	req.Args, _ = structpb.NewStruct(map[string]interface{}{"question": "world"})
	resAny, err := c.ChatEnquire(context.Background(), req.ToAny())
	if err != nil {
		t.Fatal(err)
	}
	res := mrz_v1.ToTypedRes[*pb_chat_v1.Chat, *structpb.Value](resAny)
	if res.Resp != nil {
		if res.Resp.GetStringValue() != "hello world" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
