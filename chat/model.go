package chat

import "context"

type Model interface {
	Create(ctx context.Context, params *Params) (*Response, error)
}
