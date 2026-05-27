package gollum

import (
	"context"
	"fmt"
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
	if len(p.Tools) > 0 {
		return nil, nil, fmt.Errorf("tools are not supported in Parse function")
	}
	var tools []Tool
	name := "parse_tool"
	desc := "A tool for parsing the model's response into a structured format. The model should use this tool to return the final output in a structured way."
	tools = append(tools, Tool{
		Name:        name,
		Description: desc,
		Parameters:  internal.JsonSchema[T](name, desc),
	})
	p.Tools = tools
	p.ToolChoice = &ToolChoice{
		Type: "required",
		Name: name,
	}
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
