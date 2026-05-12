package chat

import "github.com/sdkim96/gollum/chat/part"

type Usage struct {
	InputTokens  int
	OutputTokens int
	TotalTokens  int
}

type Response struct {
	Parts      []part.Part
	Usage      Usage
	StopReason string
	Model      string
}
