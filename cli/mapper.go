package cli

var MainCommandMap = CommandMap{
	"exit":   Exit,
	"help":   Help,
	"encode": Encode,
	"decode": Decode,
	"paste":  Paste,
	"clear":  Clear,
	"c":      Clear,
}
