package message

import "github.com/sdkim96/gollum/chat/part"

type ModelMessage struct {
	Parts []part.Part
}

func NewModelMessage(parts []part.Part) ModelMessage {
	return ModelMessage{
		Parts: parts,
	}
}
