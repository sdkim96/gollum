package gollum

import (
	"context"

	"github.com/sdkim96/gollum/chat"
	"github.com/sdkim96/gollum/providers/openai"
)

func main() {

	client := openai.NewClient(nil)
	gollumOpenAI := openai.NewGollumOpenAI(client)
	chat.Create(
		context.Background(),
		gollumOpenAI,
		"Write a poem about gollum",
		chat.WithInstruction("You are a helpful assistant."),
		chat.WithTemperature(0.7),
	)
	embed.Create(context.Background(), textEmbedding3Small, "Generate an embedding for the given text.")
}
