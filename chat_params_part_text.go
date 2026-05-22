package gollum

type TextPart struct {
	Content string
}

func (t TextPart) Type() string {
	return PartTypeText
}

func (t TextPart) Text() string {
	return t.Content
}

func NewTextPart(content string) TextPart {
	return TextPart{
		Content: content,
	}
}
