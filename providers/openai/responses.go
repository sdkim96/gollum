package openai

import "context"

type ResponsesService struct {
	c *Client
}

func (s *ResponsesService) Create(ctx context.Context, params *ResponsesParams) (*ResponsesResponse, error) {
	return &ResponsesResponse{}, nil
}
