package gollum

const MessageTypeModel = "model"
const MessageTypeUser = "user"

type Message interface {
	Type() string
	Parts() []Part
}

type ModelMessage struct {
	parts []Part
}

func NewModelMessage(parts []Part) ModelMessage {
	return ModelMessage{
		parts: parts,
	}
}

func (m ModelMessage) Type() string  { return MessageTypeModel }
func (m ModelMessage) Parts() []Part { return m.parts }

type UserMessage struct {
	parts []Part
}

func NewUserMessage(parts []Part) UserMessage {
	return UserMessage{
		parts: parts,
	}
}

func (u UserMessage) Type() string  { return MessageTypeUser }
func (u UserMessage) Parts() []Part { return u.parts }
