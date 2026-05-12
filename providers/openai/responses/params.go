package responses

type Params struct {
	Model        string    `json:"model"`
	Instructions *string   `json:"instructions,omitempty"`
	Input        []Message `json:"input"`
}

type Message struct {
	Role    string `json:"role"`
	Content []Part `json:"content"`
}

type Part struct {
	Type     string  `json:"type"`
	Text     *string `json:"text,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
}
