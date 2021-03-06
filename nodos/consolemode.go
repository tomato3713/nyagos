package nodos

type ModeOp interface {
	Op(mode uint32) uint32
}

type ModeReset uint32

func (ms ModeReset) Op(mode uint32) uint32 {
	return mode &^ uint32(ms)
}

type ModeSet uint32

func (ms ModeSet) Op(mode uint32) uint32 {
	return mode | uint32(ms)
}

func ChangeConsoleMode(console Handle, ops ...ModeOp) (func(), error) {
	return changeConsoleMode(console, ops...)
}

// EnableProcessInput enables Ctrl-C's signal and console's echo back.
func EnableProcessInput() (func(), error) {
	return enableProcessInput()
}
