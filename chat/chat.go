package chat

import "context"

type Creator interface {
	Create(ctx context.Context, params *Params) (*Response, error)
}

func Create(
	ctx context.Context,
	cr Creator,
	prompt string,
	options ...OptionFunc,
) (*Response, error) {

	params := &Params{}
	for _, o := range options {
		o(params)
	}
	return cr.Create(ctx, params)
}
