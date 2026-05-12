package openai

import (
	"context"

	"github.com/sdkim96/gollum/chat"
)

type GollumOpenAI struct {
	c *Client
}

func NewGollumOpenAI(client *Client) *GollumOpenAI {
	return &GollumOpenAI{
		c: client,
	}
}

func (g *GollumOpenAI) Create(ctx context.Context, params *chat.Params) (*chat.Response, error) {
	openAIParams := convertGollumParams(params)
	openaiResp, err := g.c.Responses.Create(ctx, openAIParams)
	if err != nil {
		return nil, err
	}
	return convertOpenAIResponse(openaiResp), nil
}
