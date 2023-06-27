package chat

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
	gogpt "github.com/sashabaranov/go-openai"
	"github.com/tinkler/mqttadmin/errz"
	"github.com/tinkler/mqttadmin/pkg/gs"
	"github.com/tinkler/mqttadmin/pkg/jsonz/sjson"
	"github.com/tinkler/mqttadmin/pkg/logger"
	"github.com/tinkler/ssms/backend/service/proxy/internal/gpt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Chat struct {
	Messages  []*ChatCompletionMessage `json:"messages"`
	MaxTokens int
	Model     string
	Stoped    bool
	Functions []ChatCompletionFunction
}

type ChatCompletionMessageFunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ChatCompletionMessage struct {
	Role         string                             `json:"role,omitempty"`
	Content      string                             `json:"content,omitempty"`
	Name         string                             `json:"name,omitempty"`
	FunctionCall *ChatCompletionMessageFunctionCall `json:"function_call,omitempty"`
	FinishReason string                             `json:"finish_reason,omitempty"`
}

type ChatCompletionFunctionParametersProperty struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

type ChatCompletionFunctionParameters struct {
	Type       string                                               `json:"type"`
	Properties map[string]*ChatCompletionFunctionParametersProperty `json:"properties"`
	Required   []string                                             `json:"required"`
}

type ChatCompletionFunction struct {
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Parameters  *ChatCompletionFunctionParameters `json:"parameters"`
}

func (m *Chat) Enquire(ctx context.Context, question string) (string, error) {
	if len(m.Messages) == 0 {
		return "", errz.ErrVdM("messages", "no messages", "没有消息")
	}

	request := gogpt.ChatCompletionRequest{}
	if m.Messages[0].Role != "system" {
		request.Messages = []gogpt.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "你是一个AI助手，我需要你模拟一名温柔贴心的女朋友来回答我的问题。",
			},
		}
	}
	for _, message := range m.Messages {
		request.Messages = append(request.Messages, gogpt.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
			Name:    message.Name,
		})
	}
	request.MaxTokens = m.MaxTokens
	request.Model = m.Model
	resp, err := gpt.GetClient().CreateChatCompletion(ctx, request)
	if err != nil {
		return "", errz.ErrInternal(err)
	}
	if len(resp.Choices) == 0 {
		return "", errz.ErrVdM("choices", "no choices", "没有选择")
	}

	m.Messages = append(m.Messages, &ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
		Name:    resp.Choices[0].Message.Name,
		FunctionCall: &ChatCompletionMessageFunctionCall{
			Name:      resp.Choices[0].Message.FunctionCall.Name,
			Arguments: resp.Choices[0].Message.FunctionCall.Arguments,
		},
	})
	return resp.Choices[0].Message.Content, nil
}

func (m *Chat) EnquireStream(stream gs.NullStream) error {
	req := gogpt.ChatCompletionRequest{
		Model:     gogpt.GPT4,
		MaxTokens: 20,
		Stream:    true,
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF || status.Code(err) == codes.Canceled {
			return nil
		}
		m.Stoped = false
		if len(m.Messages) == 0 {
			time.Sleep(time.Second)
			continue
		}
		if m.Model != "" {
			req.Model = m.Model
		}
		if m.MaxTokens != 0 {
			req.MaxTokens = m.MaxTokens
		}
		if len(m.Functions) > 0 {
			fs := make([]gogpt.ChatCompletionFunction, len(m.Functions))
			for i, f := range m.Functions {
				b, _ := sjson.Marshal(f)
				_ = sjson.Unmarshal(b, &fs[i])
			}
			req.Functions = fs
		}
		chatMessage := make([]gogpt.ChatCompletionMessage, len(m.Messages))
		for i, message := range m.Messages {
			chatMessage[i] = gogpt.ChatCompletionMessage{
				Role:    message.Role,
				Name:    message.Name,
				Content: message.Content,
			}
		}
		req.Messages = chatMessage
		chatStream, err := gpt.GetClient().CreateChatCompletionStream(stream.Context(), req)
		if err != nil {
			if errors.Is(context.Canceled, err) {
				logger.Info("context cancel")
			} else {
				logger.Error(err)
			}
			return status.New(codes.Internal, "create chat error").Err()
		}
		logger.Info("created chat stream")
		currentIndex := len(m.Messages)
		m.Messages = append(m.Messages, &ChatCompletionMessage{})
		for {
			response, err := chatStream.Recv()
			if errors.Is(err, io.EOF) {
				m.Stoped = true
				err = stream.Send(nil)
				if err == io.EOF || status.Code(err) == codes.Canceled {
					return nil
				}
				if err != nil {
					logger.Error(err)
				}
				break
			}
			if err != nil {
				logger.Error(err)
				m.Stoped = true
				err = stream.Send(nil)
				if err == io.EOF || status.Code(err) == codes.Canceled {
					return nil
				}
				if err != nil {
					logger.Error(err)
				}
				break
			}
			m.Messages[currentIndex].Role += response.Choices[0].Delta.Role
			m.Messages[currentIndex].FinishReason += response.Choices[0].FinishReason
			m.Messages[currentIndex].Content += response.Choices[0].Delta.Content
			logger.Info("%+v", response.Choices[0])
			if response.Choices[0].Delta.FunctionCall != nil {
				if m.Messages[currentIndex].FunctionCall == nil {
					m.Messages[currentIndex].FunctionCall = &ChatCompletionMessageFunctionCall{
						Name:      response.Choices[0].Delta.FunctionCall.Name,
						Arguments: response.Choices[0].Delta.FunctionCall.Arguments,
					}
				} else {
					m.Messages[currentIndex].FunctionCall.Name += response.Choices[0].Delta.FunctionCall.Name
					m.Messages[currentIndex].FunctionCall.Arguments += response.Choices[0].Delta.FunctionCall.Arguments
				}
			}
			err = stream.Send(nil)
			if err == io.EOF || status.Code(err) == codes.Canceled {
				return nil
			}
			if err != nil {
				logger.Error(err)
			}
		}
		chatStream.Close()
	}
}
