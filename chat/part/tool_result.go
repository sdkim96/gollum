package part

type ToolResultPart struct {
	ToolUseID string
	Content   string
}

func (t ToolResultPart) Type() string {
	return "toolResult"
}

func NewToolResultPart(toolUseID, content string) ToolResultPart {
	return ToolResultPart{
		ToolUseID: toolUseID,
		Content:   content,
	}
}
