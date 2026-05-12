package responses

import (
	"context"
	"net/http"
)

type ResponsesAPI struct {
	httpClient *http.Client
	headers    map[string]string
	baseURL    string
}

func Create(ctx context.Context, api *ResponsesAPI) {}
