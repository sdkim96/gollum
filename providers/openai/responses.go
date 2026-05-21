package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"

	"github.com/sdkim96/gollum/internal"
	oai "github.com/sdkim96/gollum/providers/openai/internal"
)

type ResponsesService struct {
	c *Client
}

func (s *ResponsesService) create(ctx context.Context, params *responsesParams) (*responsesResponse, error) {
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

	var resp responsesResponse
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ResponsesService) stream(ctx context.Context, params *responsesParams) iter.Seq2[*responsesResponse, error] {
	return func(yield func(*responsesResponse, error) bool) {
		data, err := json.Marshal(params)
		if err != nil {
			yield(nil, err)
			return
		}
		rawReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/responses", s.c.BaseURL.Host), bytes.NewReader(data))
		if err != nil {
			yield(nil, err)
			return
		}
		rawResp, err := s.c.client.Do(rawReq)
		if err != nil {
			yield(nil, err)
			return
		}
		defer rawResp.Body.Close()

		ln := internal.NewSSE(rawResp.Body, oai.NextFunc, oai.ScanFunc)
		for ln.Next() {
			var event, data string
			ln.Scan(&event, &data)
		}

		if ln.Err() != nil {
			yield(nil, ln.Err())
			return
		}
		if !yield(nil, nil) {
			return
		}
	}
}
