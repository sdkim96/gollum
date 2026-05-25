package openai

import (
	"encoding/json"
	"fmt"

	"github.com/sdkim96/gollum"
)

func toResponsesParams(p *gollum.ChatParams, stream bool) *responsesParams {

	var input []responsesInputItem
	for _, m := range p.Messages {
		input = append(input, toInputItems(m)...)
	}

	return &responsesParams{
		Model:           p.Model,
		Input:           input,
		Instructions:    p.Instruction,
		MaxOutputTokens: p.MaxOutputTokens,
		Temperature:     p.Temperature,
		Tools:           toResponsesTools(p.Tools),
		ToolChoice:      toResponsesToolChoice(p.ToolChoice),
		Reasoning:       toResponsesReasoning(p.Thinking),
		Stream:          stream,
	}
}

func toResponsesReasoning(t *gollum.Thinking) *responsesReasoning {
	if t == nil {
		return nil
	}
	return &responsesReasoning{
		Effort: t.Effort,
	}

}

func toInputItems(m gollum.Message) []responsesInputItem {
	switch m.Type() {
	case gollum.MessageTypeUser:
		var contentParts []responsesPart
		var items []responsesInputItem

		for _, p := range m.Parts() {
			switch p.Type() {
			case gollum.PartTypeText:
				text := p.(gollum.TextPart).Content
				contentParts = append(contentParts, responsesPart{
					Type: "input_text",
					Text: &text,
				})
			case gollum.PartTypeImageURL:
				url := p.(gollum.ImageURLPart).URL
				contentParts = append(contentParts, responsesPart{
					Type:     "input_image",
					ImageURL: &url,
				})
			case gollum.PartTypeImageBase64:
				ip := p.(gollum.ImageBase64Part)
				imageURL := fmt.Sprintf("data:%s;base64,%s", ip.MimeType, ip.Base64)
				contentParts = append(contentParts, responsesPart{
					Type:     "input_image",
					ImageURL: &imageURL,
				})
			case gollum.PartTypeToolResult:
				tp := p.(gollum.ToolResultPart)
				callID := tp.ToolUseID
				output := tp.Content
				items = append(items, responsesInputItem{
					Type:   "function_call_output",
					CallID: &callID,
					Output: &output,
				})
			default:
				panic(fmt.Sprintf("unsupported part type for user: %s", p.Type()))
			}
		}

		if len(contentParts) > 0 {
			items = append([]responsesInputItem{{
				Type:    "message",
				Role:    "user",
				Content: contentParts,
			}}, items...)
		}

		return items

	case gollum.MessageTypeModel:
		var parts []responsesPart
		for _, p := range m.Parts() {
			switch p.Type() {
			case gollum.PartTypeText:
				text := p.(gollum.TextPart).Content
				parts = append(parts, responsesPart{
					Type: "output_text",
					Text: &text,
				})
			default:
				panic(fmt.Sprintf("unsupported part type for assistant: %s", p.Type()))
			}
		}
		return []responsesInputItem{{
			Type:    "message",
			Role:    "assistant",
			Content: parts,
		}}

	default:
		panic(fmt.Sprintf("unsupported message type: %s", m.Type()))
	}
}

func toResponsesTools(tools []gollum.Tool) []responsesTool {
	if len(tools) == 0 {
		return nil
	}
	result := make([]responsesTool, len(tools))
	for i, t := range tools {
		desc := t.Description
		result[i] = responsesTool{
			Type:        "function",
			Name:        t.Name,
			Description: &desc,
			Parameters:  t.Parameters,
		}
	}
	return result
}

func toResponsesToolChoice(tc *gollum.ToolChoice) *responsesToolChoiceSpecific {
	if tc == nil || tc.Type != "function" {
		return nil
	}
	return &responsesToolChoiceSpecific{
		Type: "function",
		Name: tc.Name,
	}
}

func toChatResponse(resp *responsesResponse) *gollum.ChatResponse {
	var parts []gollum.Part

	for _, item := range resp.Output {
		switch item.Type {
		case "message":
			for _, c := range item.Content {
				switch c.Type {
				case "output_text":
					parts = append(parts, gollum.NewTextPart(c.Text))
				}
			}
		case "function_call":
			parts = append(parts, gollum.NewToolUsePart(item.CallID, item.Name, parseArgs(item.Arguments)))
		}
	}

	var usage gollum.ChatUsage
	if resp.Usage != nil {
		usage = gollum.ChatUsage{
			InputTokens:  resp.Usage.InputTokens,
			OutputTokens: resp.Usage.OutputTokens,
			TotalTokens:  resp.Usage.TotalTokens,
		}
	}

	return &gollum.ChatResponse{
		Message:    gollum.NewModelMessage(parts),
		Usage:      usage,
		StopReason: resp.Status,
		Model:      resp.Model,
	}
}

func parseArgs(raw string) map[string]any {
	if raw == "" {
		return nil
	}
	var m map[string]any
	json.Unmarshal([]byte(raw), &m)
	return m
}

// toStreamChatResponse converts a single SSE event into a ChatResponse.
// Returns (nil, nil) for events that don't produce a ChatResponse.
// Returns (nil, error) for error events.
func toStreamChatResponse(ev *responsesStreamEvent) (*gollum.ChatResponse, error) {
	switch ev.Type {
	case eventOutputTextDelta:
		return &gollum.ChatResponse{
			Message: gollum.NewModelMessage([]gollum.Part{
				gollum.NewTextPart(ev.Delta),
			}),
		}, nil

	case eventFunctionCallArgsDone:
		if ev.Item == nil {
			return nil, nil
		}
		return &gollum.ChatResponse{
			Message: gollum.NewModelMessage([]gollum.Part{
				gollum.NewToolUsePart(ev.Item.CallID, ev.Item.Name, parseArgs(ev.Item.Arguments)),
			}),
		}, nil

	case eventResponseCompleted:
		if ev.Response == nil {
			return nil, nil
		}
		return toChatResponse(ev.Response), nil

	case eventResponseFailed:
		if ev.Response != nil && ev.Response.Error != nil {
			return nil, fmt.Errorf("openai: %s: %s", ev.Response.Error.Code, ev.Response.Error.Message)
		}
		return nil, fmt.Errorf("openai: response failed")

	case eventError:
		if ev.Error != nil {
			return nil, fmt.Errorf("openai: %s: %s", ev.Error.Code, ev.Error.Message)
		}
		return nil, fmt.Errorf("openai: stream error")

	default:
		return nil, nil
	}
}
