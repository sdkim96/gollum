package examples

import (
	"context"
	"iter"

	"github.com/sdkim96/gollum"
	"github.com/sdkim96/gollum/openai"
)

func main() {

	client := openai.NewClient(nil)
	gollum.Create(
		context.Background(),
		client,
		"gpt-3.5-turbo",
		"Write a poem about gollum",
		gollum.WithInstruction("You are a helpful assistant."),
		gollum.WithTemperature(0.7),
	)
	embed.Create(context.Background(), textEmbedding3Small, "Generate an embedding for the given text.")
}

type itt struct {
	idx  int
	data int
}

func iterate(arr []int) iter.Seq[itt] {
	return func(yield func(itt) bool) {
		for i, v := range arr {
			if !yield(itt{idx: i, data: v}) {
				return
			}
		}
	}
}
func g() {
	arr := []int{1, 2, 3, 4, 5}
	for it := range iterate(arr) {
		println(it.idx, it.data)
	}
}
