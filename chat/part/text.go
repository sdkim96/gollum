package part

type TextPart struct {
	Content string
}

func (t TextPart) Type() string {
	return "text"
}

func NewTextPart(content string) TextPart {
	return TextPart{
		Content: content,
	}
}
