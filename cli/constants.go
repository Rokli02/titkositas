package cli

const (
	escape_reset_cursor      = "\x1B[0;0H"
	escape_clear_terminal    = "\x1B[2J"
	escape_clear_line        = "\x1B[2K\x1B[1G"
	escape_set_command_color = "\x1B[38;2;70;89;128m"
	escape_set_user_color    = "\x1B[38;2;128;128;128m"
	escape_reset_color       = "\x1B[39m"
	escape_cursor_no_blink   = "\x1B[25m"
	escape_cursor_slow_blink = "\x1B[5m"
	escape_cursor_fast_blink = "\x1B[6m"
)

const (
	Ctrl_A           = 1
	Ctrl_B           = 2
	Ctrl_C           = 3
	Ctrl_V           = 22
	Ctrl_X           = 24
	Backspace        = 8
	Enter            = 13
	Space            = 32
	Exclamation_mark = 33
	Plus             = 43
	Comma            = 44
	Minus            = 45
	Dot              = 46
	Slash            = 47
	Backslash        = 92
	Underscore       = 95
	Question_mark    = 63
	Right_click      = 100
)
