package chat

type OptionFunc func(*Params)

func WithInstruction(instruction string) OptionFunc {
	return func(o *Params) {
		o.Instruction = instruction
	}
}
func WithTemperature(temp float64) OptionFunc {
	return func(o *Params) {
		o.Temperature = temp
	}
}
