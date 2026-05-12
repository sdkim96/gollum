package openai

import (
	"context"
	"net/http"

	"github.com/sdkim96/gollum/chat"
	"github.com/sdkim96/gollum/chat/providers/openai/responses"
	"github.com/sdkim96/gollum/internal"
)

type ChatModel struct {
	model      string
	httpClient *http.Client
	invocation
}

type invocation struct {
	baseURL  string
	endpoint string
	headers  map[string]string
}

func (i *invocation) URL() string {
	return i.baseURL + i.endpoint
}

func (c *Client) NewChatModel(model string) *ChatModel {
	return &ChatModel{
		model:      model,
		httpClient: c.httpClient,
		invocation: invocation{
			baseURL:  openAIBaseURL,
			endpoint: "/responses",
			headers: internal.ImmutableCopy(map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + c.apiKey,
				"User-Agent":    "Gollum-OpenAI-Client/1.0",
			}),
		},
	}
}

func (cm *ChatModel) Create(ctx context.Context, params *chat.Params) (*chat.Response, error) {
	oairesp, err := responses.Create(ctx, cm, responses.Param(params))
	if err != nil {
		return nil, err
	}
	return buildResponse(oairesp), nil
}
