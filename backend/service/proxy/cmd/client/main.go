package main

import (
	"context"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
	mrz_v1 "github.com/tinkler/mqttadmin/mrz/v1"
	"github.com/tinkler/mqttadmin/pkg/logger"
	pb_chat_v1 "github.com/tinkler/ssms/backend/service/proxy/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	logger.ConsoleLevel = logger.LL_DEBUG
	creds, err := credentials.NewClientTLSFromFile("../../../tester/cert/ca_cert.pem", "x.test.example.com")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("35.87.142.223:5801", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb_chat_v1.NewChatGsrvClient(conn)
	stream(c)
}

func normal(c pb_chat_v1.ChatGsrvClient) {
	data := &pb_chat_v1.Chat{
		Messages: []*pb_chat_v1.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "你好,什么是Golang?",
			},
		},
		MaxTokens: 100,
		Model:     openai.GPT4,
	}
	req := mrz_v1.NewTypedModel[*pb_chat_v1.Chat, *structpb.Struct]()
	req.Data = data
	anyRes, err := c.ChatEnquire(context.Background(), req.ToAny())
	if err != nil {
		logger.Error(err)
		return
	}
	res := mrz_v1.ToTypedRes[*pb_chat_v1.Chat, *anypb.Any](anyRes)
	logger.Info(res.Data.Messages)
}

func stream(c pb_chat_v1.ChatGsrvClient) {
	data := &pb_chat_v1.Chat{
		Messages: []*pb_chat_v1.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "今天天气如何?",
			},
		},
		MaxTokens: 100,
		Model:     "gpt-3.5-turbo-0613",
		Functions: []*pb_chat_v1.ChatCompletionFunction{
			{
				Name:        "get_current_weather",
				Description: "Get the current weather in a given location",
				Parameters: &pb_chat_v1.ChatCompletionFunctionParameters{
					Type: "object",
					Properties: map[string]*pb_chat_v1.ChatCompletionFunctionParametersProperty{
						"location": {
							Type:        "string",
							Description: "The city and state, e.g. San Francisco, CA",
						},
						"unit": {
							Type: "string",
							Enum: []string{"celsius", "fahrenheit"},
						},
					},
					Required: []string{"location"},
				},
			},
		},
	}
	stream, err := c.ChatEnquireStream(context.Background())
	if err != nil {
		panic(err)
	}
	cc := make(chan struct{})
	go func() {
		defer close(cc)
		for {
			resAny, err := stream.Recv()
			if err == io.EOF || status.Code(err) == codes.Canceled {
				break
			}
			if err != nil {
				logger.Error(err)
				return
			}
			res := mrz_v1.ToDataTypedRes[*pb_chat_v1.Chat](resAny)
			data.Messages = res.Data.Messages
			msgs := make([]string, len(data.Messages))
			for i, msg := range data.Messages {
				msgs[i] = msg.Content
				if msg.FunctionCall != nil {
					fmt.Printf("%+v", msg.FunctionCall)
				}
			}
			if res.Data.Stoped {
				lastMsg := data.Messages[len(data.Messages)-1]
				logger.Info("%+v", lastMsg)
				if lastMsg.FinishReason == "function_call" {
					if lastMsg.FunctionCall == nil {
						logger.Error("function call is nil but finish reason is function_call")
					} else {
						data.Messages = append(data.Messages, &pb_chat_v1.ChatCompletionMessage{
							Role:    "function",
							Name:    lastMsg.FunctionCall.Name,
							Content: "{\"temperature\": 22, \"unit\": \"celsius\", \"description\": \"Sunny\"}",
						})
						req := mrz_v1.NewDataTypedModel[*pb_chat_v1.Chat]()
						req.Data = data
						err = stream.Send(req.ToAny())
						if err == io.EOF || status.Code(err) == codes.Canceled {
							return
						}
						if err != nil {
							logger.Error(err)
						}
					}
				}
			}
			fmt.Printf("current %t", res.Data.Stoped)
		}
	}()
	req := mrz_v1.NewDataTypedModel[*pb_chat_v1.Chat]()
	req.Data = data
	err = stream.Send(req.ToAny())
	if err == io.EOF || status.Code(err) == codes.Canceled {
		return
	}
	if err != nil {
		logger.Error(err)
	}
	<-cc
}
