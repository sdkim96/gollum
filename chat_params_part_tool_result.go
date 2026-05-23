package gollum

type ToolResultPart struct {
	ToolUseID string
	Content   string
}

func (t ToolResultPart) Type() string {
	return PartTypeToolResult
}

func (t ToolResultPart) Text() string {
	return t.Content
}

func NewToolResultPart(toolUseID, content string) ToolResultPart {
	return ToolResultPart{
		ToolUseID: toolUseID,
		Content:   content,
	}
}
