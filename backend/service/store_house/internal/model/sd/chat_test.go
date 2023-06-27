package sd

import (
	"context"
	"testing"
)

type ChatAskStream struct {
	send func(v *ChatCompletionMessage)
}

func (s *ChatAskStream) Send(v *ChatCompletionMessage) error {
	s.send(v)
	return nil
}

func (s *ChatAskStream) Recv() (*ChatCompletionMessage, error) {
	return &ChatCompletionMessage{
		Content: "你是？",
	}, nil
}

func (s *ChatAskStream) Context() context.Context {
	return context.Background()
}

func TestChatAsk(t *testing.T) {
	initTest()
	m := new(Chat)
	err := m.Ask(&ChatAskStream{
		func(v *ChatCompletionMessage) {
			t.Log(v.Content)
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
