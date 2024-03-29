package cli

type Command uint16
type CommandMap map[string]Command

const (
	UnknownCommand Command = iota
	Exit
	Help
	Encode
	Decode
	Paste
	Clear
)

func (c Command) ToString() string {
	switch c {
	case Exit:
		return "Exit"
	case Help:
		return "Help"
	case Encode:
		return "Encode"
	case Decode:
		return "Decode"
	case Paste:
		return "Paste"
	default:
		return "Unknown"
	}
}
