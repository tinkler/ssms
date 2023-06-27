package chatting

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"sync"

	mrz_v1 "github.com/tinkler/mqttadmin/mrz/v1"
	"github.com/tinkler/mqttadmin/pkg/jsonz/sjson"
	pb_chat_v1 "github.com/tinkler/ssms/backend/service/proxy/chat/v1"
	"github.com/tinkler/ssms/backend/service/proxy/pkg/model/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type chatConfig struct {
	creds        credentials.TransportCredentials
	proxyAddress string
}

type chatInst struct {
	cnf    chatConfig
	client pb_chat_v1.ChatGsrvClient
	conn   *grpc.ClientConn
}

var (
	inst     *chatInst
	instOnce sync.Once
)

func Chat() *chatInst {
	instOnce.Do(func() {
		creds, err := credentials.NewClientTLSFromFile("./ca_cert.pem", "x.test.example.com")
		if err != nil {
			panic(err)
		}
		address := os.Getenv("PROXY_ADDRESS")
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}
		inst = &chatInst{
			cnf: chatConfig{
				creds:        creds,
				proxyAddress: address,
			},
			conn:   conn,
			client: pb_chat_v1.NewChatGsrvClient(conn),
		}
	})

	return inst
}

func Close() error {
	if inst.conn != nil {
		return inst.conn.Close()
	}
	return nil
}

type stream struct {
	chat.Chat
	s pb_chat_v1.ChatGsrv_ChatEnquireStreamClient
}

func EnquireStream(ctx context.Context) (*stream, error) {
	s, err := Chat().client.ChatEnquireStream(ctx)
	if err != nil {
		return nil, err
	}
	return &stream{s: s}, nil
}

func (s *stream) Recv() (bool, error) {
	resAny, err := s.s.Recv()
	if err != nil {
		return false, err
	}
	res := mrz_v1.ToDataTypedRes[*pb_chat_v1.Chat](resAny)
	msgs := make([]*chat.ChatCompletionMessage, len(res.Data.Messages))
	for i, msg := range res.Data.Messages {
		msgs[i] = &chat.ChatCompletionMessage{
			Role:         msg.Role,
			Name:         msg.Name,
			Content:      msg.Content,
			FinishReason: msg.FinishReason,
		}
		if msg.FunctionCall != nil {
			msgs[i].FunctionCall = &chat.ChatCompletionMessageFunctionCall{
				Name:      msg.FunctionCall.Name,
				Arguments: msg.FunctionCall.Arguments,
			}
		}
	}
	s.Messages = msgs
	s.Stoped = res.Data.Stoped
	s.Model = res.Data.Model
	funs := make([]*chat.ChatCompletionFunction, len(res.Data.Functions))
	for i, f := range res.Data.Functions {
		funs[i] = &chat.ChatCompletionFunction{
			Name:        f.Name,
			Description: f.Description,
		}
		if f.Parameters != nil {
			funs[i].Parameters = &chat.ChatCompletionFunctionParameters{
				Required:   f.Parameters.Required,
				Properties: make(map[string]*chat.ChatCompletionFunctionParametersProperty),
			}
			for k, v := range f.Parameters.Properties {
				funs[i].Parameters.Properties[k] = &chat.ChatCompletionFunctionParametersProperty{
					Type:        v.Type,
					Description: v.Description,
				}
			}
		}
	}
	return true, nil
}

func (s *stream) Send(_ *struct{}) error {
	req := mrz_v1.NewDataTypedModel[*pb_chat_v1.Chat]()
	req.Data = new(pb_chat_v1.Chat)
	byt, _ := sjson.Marshal(s.Chat)
	err := sjson.Unmarshal(byt, req.Data)
	if err != nil {
		return err
	}
	return s.s.Send(req.ToAny())
}

func (s *stream) CloseSend() error {
	return s.s.CloseSend()
}

func (s *stream) Last() *chat.ChatCompletionMessage {
	return s.Messages[len(s.Messages)-1]
}

func (s *stream) AddFunctionMessage(name string, content any, description string) {
	kv := make(map[string]interface{})
	switch c := content.(type) {
	case string:
		if len(c) > 0 {
			_ = json.Unmarshal([]byte(c), &kv)
		}
	case []byte:
		if len(c) > 0 {
			_ = json.Unmarshal(c, &kv)
		}
	case nil:
	default:
		rt := reflect.TypeOf(content)
		if rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}
		if rt.Kind() != reflect.Struct && rt.Kind() != reflect.Array && rt.Kind() != reflect.Slice {
			panic("unsupported type")
		}
		byt, _ := json.Marshal(content)
		_ = json.Unmarshal(byt, &kv)
	}
	if description != "" {
		kv["description"] = description
	}

	result, _ := json.Marshal(kv)

	msg := chat.ChatCompletionMessage{
		Role:    "function",
		Name:    name,
		Content: string(result),
	}
	s.Messages = append(s.Messages, &msg)
}

func (s *stream) AddUserMessage(content string) {
	msg := chat.ChatCompletionMessage{
		Role:    "user",
		Content: content,
	}
	s.Messages = append(s.Messages, &msg)
}
