package gollum

import "encoding/json"

type ChatParams struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Instruction *string   `json:"instruction,omitempty"`

	MaxOutputTokens *int     `json:"max_output_tokens,omitempty"`
	Temperature     *float64 `json:"temperature,omitempty"`
	TopP            *float64 `json:"top_p,omitempty"`
	StopSequences   []string `json:"stop_sequences,omitempty"`

	Tools      []Tool      `json:"tools,omitempty"`
	ToolChoice *ToolChoice `json:"tool_choice,omitempty"`
}

type Tool struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Parameters  json.RawMessage `json:"parameters"`
}

type ToolChoice struct {
	Type string `json:"type"` // "auto" | "required" | "none" | "function"
	Name string `json:"name,omitempty"`
}
