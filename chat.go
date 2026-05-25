package gollum

import (
	"context"
	"iter"

	"github.com/sdkim96/gollum/internal"
)

type Creator interface {
	Create(ctx context.Context, p *ChatParams) (*ChatResponse, error)
}

type Streamer interface {
	Stream(ctx context.Context, p *ChatParams) iter.Seq2[*ChatResponse, error]
}

func Create(ctx context.Context, c Creator, p *ChatParams) (*ChatResponse, error) {
	return c.Create(ctx, p)
}

func Stream(ctx context.Context, s Streamer, p *ChatParams) iter.Seq2[*ChatResponse, error] {
	return s.Stream(ctx, p)
}

func Parse[T any](ctx context.Context, c Creator, p *ChatParams) (*T, *ChatResponse, error) {
	resp, err := c.Create(ctx, p)
	if err != nil {
		return nil, nil, err
	}

	result, err := internal.Parse[T]([]byte(resp.Text()))
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
