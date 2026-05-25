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
		for ev, err := range c.Responses.stream(ctx, toResponsesParams(gp, true)) {
			if err != nil {
				yield(nil, err)
				return
			}
			resp, err := toStreamChatResponse(ev)
			if err != nil {
				yield(nil, err)
				return
			}
			if resp != nil {
				if !yield(resp, nil) {
					return
				}
			}
		}
	}
}

func inspectModel(model string) error {
	if !slices.Contains(allowedModels, model) {
		return fmt.Errorf("unsupported model: %s", model)
	}
	return nil
}
