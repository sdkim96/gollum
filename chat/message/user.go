package message

import "github.com/sdkim96/gollum/chat/part"

type UserMessage struct {
	Parts []part.Part
}

func NewUserMessage(parts []part.Part) UserMessage {
	return UserMessage{
		Parts: parts,
	}
}
