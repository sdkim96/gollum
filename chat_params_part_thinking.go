package gollum

type ThinkingPart struct {
	Content string
}

func (t ThinkingPart) Type() string {
	return PartTypeThinking
}

func NewThinkingPart(content string) ThinkingPart {
	return ThinkingPart{
		Content: content,
	}
}
