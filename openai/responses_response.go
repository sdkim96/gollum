package openai


// responsesResponse is the body returned by POST /v1/responses.
// Converted to chat.Response via toChatResponse; never exposed to users.
type responsesResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"` // always "response"
	CreatedAt int64  `json:"created_at"`
	Model     string `json:"model"`
	Status    string `json:"status"` // "completed" | "in_progress" | "failed" | "incomplete"

	Output []responsesOutputItem `json:"output"`
	Usage  *responsesUsage       `json:"usage,omitempty"`
	Error  *responsesError       `json:"error,omitempty"`
}

// responsesOutputItem is an element of the output array (a union).
// Field usage depends on Type:
//   - "message"       → Role + Content (assistant reply)
//   - "function_call" → CallID + Name + Arguments (becomes chat.ToolUsePart)
//   - "reasoning"     → Summary + EncryptedContent (becomes chat.ThinkingPart)
type responsesOutputItem struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status,omitempty"`

	// for "message"
	Role    string                `json:"role,omitempty"` // "assistant"
	Content []responsesOutputPart `json:"content,omitempty"`

	// for "function_call"
	CallID    string `json:"call_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"` // raw JSON string

	// for "reasoning"
	Summary          []responsesReasoningSummary `json:"summary,omitempty"`
	EncryptedContent string                      `json:"encrypted_content,omitempty"`
}

// responsesOutputPart is a content element inside an assistant message.
// Field usage depends on Type:
//   - "output_text" → Text
//   - "refusal"     → Refusal
//
// Annotations (file_citation, url_citation, etc.) are intentionally omitted
// until gollum needs to surface them.
type responsesOutputPart struct {
	Type    string `json:"type"`
	Text    string `json:"text,omitempty"`
	Refusal string `json:"refusal,omitempty"`
}

// responsesReasoningSummary is a chain-of-thought summary block from
// reasoning-capable models (gpt-5, o-series).
type responsesReasoningSummary struct {
	Type string `json:"type"` // "summary_text"
	Text string `json:"text"`
}

// responsesUsage holds token accounting for the request.
type responsesUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
	TotalTokens  int `json:"total_tokens"`

	InputTokensDetails  *responsesInputTokenDetails  `json:"input_tokens_details,omitempty"`
	OutputTokensDetails *responsesOutputTokenDetails `json:"output_tokens_details,omitempty"`
}

type responsesInputTokenDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type responsesOutputTokenDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type responsesError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

