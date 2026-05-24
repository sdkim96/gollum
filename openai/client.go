package openai

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {

	// HTTP client used to communicate with the OpenAI API.
	// The client should be shared where possible to take advantage of HTTP connection reuse.
	client *http.Client
	mu     sync.Mutex

	BaseURL   *url.URL
	UserAgent string

	Responses  *ResponsesService
	Embeddings *EmbeddingsService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{
		client: httpClient,
	}
	c.initialize()
	return c
}

func (c *Client) WithAPIKey(apiKey string) *Client {
	c2 := c.copy()
	defer c2.initialize()
	c2transport := c2.client.Transport
	if c2transport == nil {
		c2transport = http.DefaultTransport
	}

	c2.client.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.Header.Set(headerAuthorization, fmt.Sprintf("Bearer %s", apiKey))
			return c2transport.RoundTrip(req)
		},
	)
	return c2
}

func (c *Client) WithHooks(bef func(req *http.Request), after func(req *http.Request, resp *http.Response, err error)) *Client {
	c2 := c.copy()
	defer c2.initialize()
	c2transport := c2.client.Transport
	if c2transport == nil {
		c2transport = http.DefaultTransport
	}

	c2.client.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			bef(req)
			resp, err := c2transport.RoundTrip(req)
			after(req, resp, err)
			return resp, err
		},
	)
	return c2
}

// initialize guarantees default values and the services for the Client struct.
func (c *Client) initialize() {
	if c.BaseURL == nil {
		baseURL, _ := url.Parse(defaultBaseURL)
		c.BaseURL = baseURL
	}
	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}
	c.Responses = &ResponsesService{c: c}
	c.Embeddings = &EmbeddingsService{c: c}
}

// copy returns a shallow copy of the Client. A new *http.Client is allocated,
// but its fields (Transport, Jar, CheckRedirect) reference the same objects
// as the original — this preserves connection pool sharing while allowing
// the copy to swap its Transport independently of the original.
func (c *Client) copy() *Client {
	c.mu.Lock()
	c2 := &Client{
		client:    &http.Client{},
		BaseURL:   c.BaseURL,
		UserAgent: c.UserAgent,
	}
	c.mu.Unlock()

	// Copying the pointers of configurations.
	if c.client != nil {
		c2.client.Transport = c.client.Transport
		c2.client.Timeout = c.client.Timeout
		c2.client.Jar = c.client.Jar
		c2.client.CheckRedirect = c.client.CheckRedirect
	}
	return c2
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
