package chat

import "github.com/sdkim96/gollum/chat/message"

type Params struct {
	Messages    []message.Message
	Instruction string
	Temperature float64
}
