package openai

import (
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
		Stream:          stream,
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
