package chat

import "context"

func Create(
	ctx context.Context,
	model Model,
	prompt string,
	options ...OptionFunc,
) (*Response, error) {

	params := &Params{}
	for _, o := range options {
		o(params)
	}
	return model.Create(ctx, params)
}
