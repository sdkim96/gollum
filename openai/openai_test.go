package openai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/sdkim96/gollum"
)

func loadTestData(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func realOpenAIClient() *Client {
	apiKey := os.Getenv("OPENAI_API_KEY")
	return NewClient(nil).
		WithAPIKey(apiKey).
		WithHooks(func(req *http.Request) {
			fmt.Printf("Request Starts")
		}, func(req *http.Request, resp *http.Response, err error) {
			d, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			resp.Body = io.NopCloser(bytes.NewBuffer(d))

			os.WriteFile("req.json", d, 0644)
			fmt.Printf("Request Ends")
		})
}

func newNormalChatParams() *gollum.ChatParams {
	var messages []gollum.Message
	var parts []gollum.Part

	parts = append(parts, gollum.NewTextPart("You are a helpful assistant."))
	messages = append(messages, gollum.NewUserMessage(parts))

	return &gollum.ChatParams{
		Model:    "gpt-4o-mini",
		Messages: messages,
	}
}

func newMultiturnChatParams() *gollum.ChatParams {
	var messages []gollum.Message
	var parts0 []gollum.Part
	var parts1 []gollum.Part
	var parts2 []gollum.Part

	parts0 = append(
		parts0,
		gollum.NewTextPart("My name is gollum, looking for precious."),
	)
	parts1 = append(
		parts1,
		gollum.NewTextPart("Hello gollum! What is your precious?"),
	)
	parts2 = append(
		parts2,
		gollum.NewTextPart("My precious is the One Ring."),
		gollum.NewTextPart("Do you know where it is?"),
	)
	messages = append(
		messages,
		gollum.NewUserMessage(parts0),
		gollum.NewModelMessage(parts1),
		gollum.NewUserMessage(parts2),
	)

	return &gollum.ChatParams{
		Model:    "gpt-4o-mini",
		Messages: messages,
	}

}

func TestResponsesCreate(t *testing.T) {
	c := realOpenAIClient()
	c.Create(context.Background(), newMultiturnChatParams())
}
