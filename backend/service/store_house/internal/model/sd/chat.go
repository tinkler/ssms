package sd

import (
	"encoding/json"
	"io"

	"github.com/tinkler/mqttadmin/pkg/gs"
	"github.com/tinkler/mqttadmin/pkg/logger"
	"github.com/tinkler/mqttadmin/pkg/status"
	"github.com/tinkler/ssms/backend/service/proxy/pkg/chatting"
	"github.com/tinkler/ssms/backend/service/proxy/pkg/model/chat"
	"google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

type Chat struct{}

type ChatCompletionMessage struct {
	Role         string `json:"role,omitempty"`
	Content      string `json:"content,omitempty"`
	Name         string `json:"name,omitempty"`
	FinishReason string `json:"finish_reason,omitempty"`
}

func (m *Chat) Ask(stream gs.Stream[*ChatCompletionMessage]) error {
	chatting, err := chatting.EnquireStream(stream.Context())
	if err != nil {
		return status.StatusInternalServer(err)
	}
	defer chatting.CloseSend()

	chatting.Messages = []*chat.ChatCompletionMessage{
		{
			Role:    "system",
			Content: `你是一个仓库管理系统,你叫AI进销存，可以查询商品和订单信息以及库存，只用中文回答并且只回答仓库数据相关的问题。`,
		},
	}
	chatting.MaxTokens = 200
	chatting.Model = "gpt-3.5-turbo-0613"
	chatting.Functions = []chat.ChatCompletionFunction{
		{
			Name:        "get_goods_information_by_name",
			Description: "Search for the goods information by name.",
			Parameters: &chat.ChatCompletionFunctionParameters{
				Type: "object",
				Properties: map[string]*chat.ChatCompletionFunctionParametersProperty{
					"name": {
						Type:        "string",
						Description: "The full name or part of the name, e.g. 家家面",
					},
				},
				Required: []string{"name"},
			},
		},
	}
	cc := make(chan struct{})
	go func() {
		defer close(cc)
		for {
			_, err := chatting.Recv()
			if err == io.EOF || gstatus.Code(err) == codes.Canceled {
				return
			}
			if err != nil {
				logger.Error(err)
				return
			}
			last := chatting.Last()
			if last.FunctionCall != nil {
				if !chatting.Stoped {
					continue
				}
				switch last.FunctionCall.Name {
				case "get_goods_information_by_name":
					args := make(map[string]string)
					err := json.Unmarshal([]byte(last.FunctionCall.Arguments), &args)
					if err != nil {
						logger.Error(err)
						chatting.AddFunctionMessage("get_goods_information_by_name", nil, "error caused!")
					} else {
						data, err := (new(GoodsviewCRUD)).SearchForGoodsViewByName(stream.Context(), args["name"])
						if err != nil {
							logger.Error(err)
							chatting.AddFunctionMessage("get_goods_information_by_name", nil, "error caused!")
						} else {
							if len(data) == 0 {
								chatting.AddFunctionMessage("get_goods_information_by_name", nil, "data not found!")
							} else if len(data) > 1 {
								var (
									currentInHand = 0.0
									index         int
								)
								for i, d := range data {
									if d.OnHand > currentInHand {
										currentInHand = d.OnHand
										index = i
									}
								}
								chatting.AddFunctionMessage("get_goods_information_by_name", data[index], "Found more than one goods. This is the most stocked goods")
							} else {
								chatting.AddFunctionMessage("get_goods_information_by_name", data[0], "Code is the 2D code of goods, on_hand is the remaining quantity, name is the full name of goods, goods_type_name is the category of goods")
							}
						}

					}
					err = chatting.Send(nil)
					if err != nil {
						logger.Error(err)
						return
					}
				}
			}
			if last.Role == "assistant" {
				err = stream.Send(&ChatCompletionMessage{
					Role:         last.Role,
					Content:      last.Content,
					Name:         last.Name,
					FinishReason: last.FinishReason,
				})
				if err == io.EOF || gstatus.Code(err) == codes.Canceled {
					return
				}
				if err != nil {
					logger.Error(err)
					return
				}
				if last.FinishReason == "stop" {
					logger.Info("%+v", chatting)
					return
				}
			}
		}
	}()

	ask, err := stream.Recv()
	if err == io.EOF || gstatus.Code(err) == codes.Canceled {
		return nil
	}
	if err != nil {
		logger.Error(err)
		return err
	}
	if ask.Content != "" {
		chatting.AddUserMessage(ask.Content)
		err := chatting.Send(nil)
		if err != nil {
			return status.StatusInternalServer(err)
		}
	}
	<-cc

	return nil
}
