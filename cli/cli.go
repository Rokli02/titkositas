package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"

	"golang.design/x/clipboard"
	"golang.org/x/term"
)

type CommandHandler struct {
	commandMap CommandMap
}

func NewCommandHandler(commandMap CommandMap) *CommandHandler {
	ch := &CommandHandler{
		commandMap: commandMap,
	}

	term.MakeRaw(int(os.Stdin.Fd()))
	clipboard.Init()

	ch.Clear()
	fmt.Printf(
		"%sType in 'help' to write out every command, that you can use!%s%s\n",
		escape_set_command_color, escape_reset_color, escape_cursor_slow_blink,
	)

	return ch
}

func (ch *CommandHandler) ReadLine() string {
	var lineBuilder []rune = make([]rune, 0)
	var byteBuffer []byte = make([]byte, 4)

	fmt.Print(escape_set_user_color, "$> ")

	defer func() {
		fmt.Println(escape_reset_color)
	}()

	for {
		os.Stdin.Read(byteBuffer)
		pressedChar, _ := utf8.DecodeRune(byteBuffer)

		switch pressedChar {
		case Ctrl_C:
			continue
		case Enter:
			goto escape_loop
		case Backspace:
			if len(lineBuilder) > 0 {
				fmt.Print("\x1B[1D \x1B[1D")
				lineBuilder = lineBuilder[:len(lineBuilder)-1]
			}

			continue
		case Ctrl_V:
			txtFromClipboard := string(clipboard.Read(clipboard.FmtText))
			lineBuilder = append(lineBuilder, []rune(txtFromClipboard)...)

			fmt.Print(txtFromClipboard)
			continue
		}

		lineBuilder = append(lineBuilder, pressedChar)
		fmt.Print(string(pressedChar))
	}

escape_loop:

	return string(lineBuilder)
}

func (ch *CommandHandler) ReadCommand() (c Command, line string) {
	line = ch.ReadLine()
	words := strings.Split(line, " ")
	size := len(words)

	if size < 1 {
		return
	}

	if command, hasCommand := ch.commandMap[words[0]]; hasCommand {
		c = command
	}

	return
}

func (cr *CommandHandler) Printf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("%s%s%s", escape_set_command_color, format, escape_reset_color), a...)
}

func (cr *CommandHandler) Println(a ...any) {
	fmt.Print(escape_set_command_color)
	fmt.Println(a...)
	fmt.Print(escape_reset_color)
}

func (cr *CommandHandler) Print(a ...any) {
	fmt.Print(escape_set_command_color)
	fmt.Print(a...)
	fmt.Print(escape_reset_color)
}

func (ch *CommandHandler) Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
