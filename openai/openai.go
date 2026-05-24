package openai

import (
	"context"
	"fmt"
	"iter"
	"slices"

	"github.com/sdkim96/gollum"
)

const (
	Version = "v0.1.0"

	headerContentType   = "Content-Type"
	headerAuthorization = "Authorization"
	headerUserAgent     = "User-Agent"

	defaultBaseURL   = "https://api.openai.com/v1"
	defaultUserAgent = "gollum/openai" + "/" + Version

	mediaTypeJSON = "application/json"
)

var allowedModels = []string{
	"gpt-4o-mini",
	"gpt-4o",
	"gpt-5-nano",
	"gpt-5-mini",
	"gpt-5",
}

func (c *Client) Create(ctx context.Context, gp *gollum.ChatParams) (*gollum.ChatResponse, error) {
	if err := inspectModel(gp.Model); err != nil {
		return nil, err
	}
	resp, err := c.Responses.create(ctx, toResponsesParams(gp, false))
	if err != nil {
		return nil, err
	}
	return toChatResponse(resp), nil
}

func (c *Client) Stream(ctx context.Context, gp *gollum.ChatParams) iter.Seq2[*gollum.ChatResponse, error] {
	return func(yield func(*gollum.ChatResponse, error) bool) {
		if err := inspectModel(gp.Model); err != nil {
			yield(nil, err)
			return
		}
		_ = toResponsesParams(gp, true)
		// for openaiResp := range c.Responses.stream(ctx, openAIParams) {
		// 	if openaiResp.Err != nil {
		// 		yield(nil, openaiResp.Err)
		// 		return
		// 	}
		// 	yield(toChatResponse(openaiResp.Value), nil)
		// }
	}
}

func inspectModel(model string) error {
	if !slices.Contains(allowedModels, model) {
		return fmt.Errorf("unsupported model: %s", model)
	}
	return nil
}
