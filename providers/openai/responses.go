package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponsesService struct {
	c *Client
}

func (s *ResponsesService) Create(ctx context.Context, params *ResponsesParams) (*ResponsesResponse, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	rawReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/responses", s.c.BaseURL.Host), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	rawResp, err := s.c.client.Do(rawReq)
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	var resp ResponsesResponse
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
