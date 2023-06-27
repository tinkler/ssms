package smart

type ChatMessage struct {
	Content string
	Sender  string
}

type SmartChat struct {
	Messages []*ChatMessage
}
