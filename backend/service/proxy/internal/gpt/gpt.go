package gpt

import (
	"os"
	"sync"

	gogpt "github.com/sashabaranov/go-openai"
)

var (
	client     *gogpt.Client
	clientOnce sync.Once
)

func GetClient() *gogpt.Client {
	clientOnce.Do(func() {
		client = gogpt.NewClientWithConfig(gogpt.DefaultConfig(os.Getenv("OPENAI_API_KEY")))
	})
	return client
}
