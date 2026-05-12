package gollum

import (
	"context"
	
	"github.com/sdkim96/gollum/resource/chat",
	"github.com/sdkim96/gollum/resource/embed",
)


func main() {

	client := openai.NewClient(
		openai.WithAPIKey("your_api_key_here"),
	)
	gpt4oMini := client.NewChatModel("gpt-4o-mini")
	textEmbedding3Small := client.NewEmbeddingModel("text-embedding-3-small")

	chat.Create(context.Background(), gpt4oMini, "Write a poem about gollum", options.WithInstruction("You are a helpful assistant."), options.WithTemperature(0.7))
	embed.Create(context.Background(), textEmbedding3Small, "Generate an embedding for the given text.")
}