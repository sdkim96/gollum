package embedding

import "context"

func Create(
	ctx context.Context,
	model model.LLMModel,
	inputs []string,
	options ...options.EmbedOptionFunc,
) (*EmbedResponse, error) {

}
