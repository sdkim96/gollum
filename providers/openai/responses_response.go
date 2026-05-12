package openai

import "github.com/sdkim96/gollum/chat"

type ResponsesResponse struct{}

func convertOpenAIResponse(resp *ResponsesResponse) *chat.Response {
	return &chat.Response{}
}
