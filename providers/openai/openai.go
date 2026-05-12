package openai

import "net/http"

const openAIBaseURL = "https://api.openai.com/v1"

type Client struct {
	apiKey     string
	httpClient *http.Client
}

type ClientOptionFunc func(*Client)

func NewClient(apiKey string, opts ...ClientOptionFunc) *Client {
	client := &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHTTPClient(c *http.Client) ClientOptionFunc {
	return func(client *Client) {
		client.httpClient = c
	}
}
