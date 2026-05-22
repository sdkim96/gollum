package gollum

const (
	PartTypeText        = "text"
	PartTypeImageURL    = "imageURL"
	PartTypeImageBase64 = "imageBase64"
	PartTypeThinking    = "thinking"
	PartTypeToolUse     = "toolUse"
	PartTypeToolResult  = "toolResult"
)

type Part interface {
	Type() string
	Text() string
}
