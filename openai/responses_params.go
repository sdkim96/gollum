package openai

import (
	"encoding/json"
)

// responsesParams is the request body for POST /v1/responses.
// gollum users work with chat.Params; this type is populated by the converter.
type responsesParams struct {
	Model string               `json:"model"`
	Input []responsesInputItem `json:"input"`

	Instructions    *string  `json:"instructions,omitempty"`
	MaxOutputTokens *int     `json:"max_output_tokens,omitempty"`
	Temperature     *float64 `json:"temperature,omitempty"`

	Tools             []responsesTool              `json:"tools,omitempty"`
	ToolChoice        *responsesToolChoiceSpecific `json:"tool_choice,omitempty"`
	ParallelToolCalls *bool                        `json:"parallel_tool_calls,omitempty"`

	Reasoning *responsesReasoning `json:"reasoning,omitempty"`

	Stream bool `json:"stream,omitempty"`
}

// responsesInputItem is an element of the input array (a union).
// Field usage depends on Type. Always construct via the helpers below.
//   - "message"              → Role + Content
//   - "function_call_output" → CallID + Output
type responsesInputItem struct {
	Type string `json:"type,omitempty"`

	// for message
	Role    string          `json:"role,omitempty"` // "user" | "assistant"
	Content []responsesPart `json:"content,omitempty"`

	// for function_call_output
	CallID *string `json:"call_id,omitempty"`
	Output *string `json:"output,omitempty"`
}

// responsesPart is a content element inside a message.
// Field usage depends on Type.
//   - "input_text"  → Text                    (user message)
//   - "output_text" → Text                    (assistant message, for stateless multi-turn)
//   - "input_image" → ImageURL (+ Detail)     (user message)
type responsesPart struct {
	Type     string  `json:"type"`
	Text     *string `json:"text,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	Detail   *string `json:"detail,omitempty"` // "low" | "high" | "auto"
}

// responsesTool defines a function tool.
type responsesTool struct {
	Type        string          `json:"type"` // "function"
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Parameters  json.RawMessage `json:"parameters"`
	Strict      *bool           `json:"strict,omitempty"`
}

// responsesToolChoiceSpecific forces a specific tool call.
type responsesToolChoiceSpecific struct {
	Type string `json:"type"` // "function"
	Name string `json:"name"`
}

// responsesReasoning configures reasoning/thinking for capable models.
type responsesReasoning struct {
	Effort  *string `json:"effort,omitempty"`  // "low" | "medium" | "high"
	Summary *string `json:"summary,omitempty"` // "auto" | "concise" | "detailed" | "none"
}
