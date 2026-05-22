package gollum

import "strings"

type ChatUsage struct {
	InputTokens  int
	OutputTokens int
	TotalTokens  int
}

type ChatResponse struct {
	Message    ModelMessage
	Usage      ChatUsage
	StopReason string
	Model      string
}

func (c *ChatResponse) Text() string {
	parts := c.Message.Parts()
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part.Text())
	}
	return sb.String()
}
