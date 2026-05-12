package part

type ThinkingPart struct {
	Content string
}

func (t ThinkingPart) Type() string {
	return "thinking"
}

func NewThinkingPart(content string) ThinkingPart {
	return ThinkingPart{
		Content: content,
	}
}
