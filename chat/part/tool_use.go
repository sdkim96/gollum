package part

type ToolUsePart struct {
	ToolUseID string
	ToolName  string
	Args      map[string]any
}

func (t ToolUsePart) Type() string {
	return "toolUse"
}

func NewToolUsePart(toolUseID, toolName string, args map[string]any) ToolUsePart {
	return ToolUsePart{
		ToolUseID: toolUseID,
		ToolName:  toolName,
		Args:      args,
	}
}
