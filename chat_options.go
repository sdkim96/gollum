package gollum

type OptionFunc func(*ChatParams)

func WithInstruction(instruction string) OptionFunc {
	return func(o *ChatParams) {
		o.Instruction = &instruction
	}
}
func WithTemperature(temp float64) OptionFunc {
	return func(o *ChatParams) {
		o.Temperature = &temp
	}
}
