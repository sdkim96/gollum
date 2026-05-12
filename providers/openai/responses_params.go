package openai

import "github.com/sdkim96/gollum/chat"

type ResponsesParams struct {
	Model        string             `json:"model"`
	Instructions *string            `json:"instructions,omitempty"`
	Input        []ResponsesMessage `json:"input"`
}

type ResponsesMessage struct {
	Role    string          `json:"role"`
	Content []ResponsesPart `json:"content"`
}

type ResponsesPart struct {
	Type     string  `json:"type"`
	Text     *string `json:"text,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
}

func convertGollumParams(params *chat.Params) *ResponsesParams {
	return &ResponsesParams{}
}
