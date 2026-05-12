package message

import "github.com/sdkim96/gollum/chat/part"

type Message interface {
	Type() string
	Parts() []part.Part
}
