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

	Thinking *Thinking `json:"thinking,omitempty"`
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

// Thinking represents the thinking configuration.
// Each field is optional and may be used to guide the model's reasoning process.
//
// OpenAI models (e.g., gpt-5o-nano) may use the Effort field,
// while Anthropic models (e.g., claude-2) may use the Budget field.
//
// The other field that the vendor does not use will be ignored.
type Thinking struct {

	// Budget field is valid for Anthropic and Google models.
	Budget *int `json:"budget,omitempty"`

	// Effort field is valid for OpenAI models.
	Effort *string `json:"effort,omitempty"` // "low" | "medium" | "high"
}
