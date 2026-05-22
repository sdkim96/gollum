package gollum

import "context"

type Embedder interface {
	Embed(ctx context.Context, params *EmbeddingParams) (*EmbeddingResponse, error)
}

func Embed(ctx context.Context, e Embedder, p *EmbeddingParams) (*EmbeddingResponse, error) {
	return e.Embed(ctx, p)
}
