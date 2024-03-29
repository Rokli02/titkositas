package encryption

var (
	// source -> target
	encoderMap map[rune][]rune
	// target -> source
	decoderMap map[rune]rune
)

var (
	maxMixTry = 3
)
