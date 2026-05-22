package openai

import (
	"fmt"

	"github.com/sdkim96/gollum"
)

func toResponsesParams(params *gollum.ChatParams, stream bool) *responsesParams {

	for _, m := range params.Messages {
		switch m.Type() {

		// assistant or tool message for openai.
		case gollum.MessageTypeModel:
			for _, ps := range m.Parts() {
				switch ps.Type() {
				case gollum.PartTypeText:
					newAssistantMessage(append([]responsesPart(nil), newOutputTextPart(ps.Text())))
				case gollum.PartTypeImage:
					newAssistantMessage(append([]responsesPart(nil), newInputImagePart(ps.ImageURL())))
				}
			}

			// user message for openai.
		case gollum.MessageTypeUser:
			for _, ps := range m.Parts() {
				switch ps.Type() {
				case gollum.PartTypeText:
					newUserMessage(append([]newInputTextPart(nil), newInputTextPart(ps.Text())))
				case gollum.PartTypeImage:
					newUserMessage(append([]responsesPart(nil), newInputImagePart(ps.ImageURL())))
				}
			}
		}
	}

	return &responsesParams{
		Model:           params.Model,
		Input:           toResponsesInput(params.Messages),
		Instructions:    params.Instruction,
		MaxOutputTokens: params.MaxOutputTokens,
		Temperature:     params.Temperature,
		Tools:           toResponsesTools(params.Tools),
		ToolChoice:      toResponsesToolChoice(params.ToolChoice),
		Stream:          stream,
	}
}

func toResponsesPart(p gollum.Part) responsesPart {
	switch p.Type() {
	case gollum.PartTypeText:
		return newInputTextPart(p.Text())
	case gollum.PartTypeImageURL:
		ip := p.(gollum.ImageURLPart)
		return newInputImagePart(fmt.Sprintf("%s", ip.URL))
	case gollum.PartTypeImageBase64:
		ip := p.(gollum.ImageBase64Part)
		return newInputImagePart(fmt.Sprintf("data:%s;base64,%s", ip.MimeType, ip.Base64))
	case gollum.PartTypeThinking:
		return newThinkingPart()
	default:
		panic("unsupported part type: " + p.Type())
	}
}
