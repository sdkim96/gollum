package openai

import (
	"context"
	"iter"

	"github.com/sdkim96/gollum/chat"
)

func (c *Client) Create(ctx context.Context, params *chat.Params) (*chat.Response, error) {
	openAIParams := toResponsesParams(params)
	openaiResp, err := c.Responses.create(ctx, openAIParams)
	if err != nil {
		return nil, err
	}
	return toChatResponse(openaiResp), nil
}

func (c *Client) Stream(ctx context.Context, params *chat.Params) iter.Seq2[*chat.Response, error] {
	openAIParams := toResponsesParams(params)
	for resp, err := range c.Responses.stream(ctx, openAIParams) {
		if err != nil {
			break
		}
	}

	return toChatResponse(openaiResp), nil
}

func (c *Client) Parse(ctx context.Context, params *chat.Params, out any) (*chat.Response, error) {

}
