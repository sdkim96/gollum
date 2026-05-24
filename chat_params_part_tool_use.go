package gollum

type ToolUsePart struct {
	ToolUseID string
	ToolName  string
	Args      map[string]any
}

func (t ToolUsePart) Type() string {
	return PartTypeToolUse
}

func (t ToolUsePart) Text() string {
	return ""
}

func NewToolUsePart(toolUseID, toolName string, args map[string]any) ToolUsePart {
	return ToolUsePart{
		ToolUseID: toolUseID,
		ToolName:  toolName,
		Args:      args,
	}
}
