package main

import (
	"strings"
	"titkositas/assets"
	"titkositas/cli"
	"titkositas/encryption"

	"golang.design/x/clipboard"
)

func main() {
	assets.LoadAssets()
	encryption.InitializeCoderMaps(assets.EncryptionTable.UseAdditionals)
	commandHandler := cli.NewCommandHandler(cli.MainCommandMap)

	for {
		command, line := commandHandler.ReadCommand()

		switch command {
		case cli.Help:
			words := strings.Split(line, " ")
			if len(words) <= 1 {
				commandHandler.Println(assets.Texts.Help)

				break
			}

			switch cli.MainCommandMap[words[1]] {
			case cli.Encode:
				commandHandler.Println(assets.Texts.Encode, encryption.GetEncoderMap())
			case cli.Decode:
				commandHandler.Println(assets.Texts.Decode, encryption.GetDecoderMap())
			default:
				commandHandler.Printf("Unknown help command!\n")
			}
		case cli.Exit:
			goto endOfMain
		case cli.Encode:
			var firstSpaceIndex int = strings.Index(line, " ")
			var hasText bool

			if firstSpaceIndex != -1 {
				afterCommandText := strings.TrimSpace(line[firstSpaceIndex:])

				if len(afterCommandText) > 0 {
					line = afterCommandText
					hasText = true
				}
			}

			if !hasText {
				commandHandler.Println("Write down a sentence, which will be encrypted:")
				command, line = commandHandler.ReadCommand()

				if command == cli.Paste {
					line = string(clipboard.Read(clipboard.FmtText))
				}
			}

			commandHandler.Println(encryption.Encode(line))
		case cli.Decode:
			var firstSpaceIndex int = strings.Index(line, " ")
			var hasText bool

			if firstSpaceIndex != -1 {
				afterCommandText := strings.TrimSpace(line[firstSpaceIndex:])

				if len(afterCommandText) > 0 {
					line = afterCommandText
					hasText = true
				}
			}

			if !hasText {
				commandHandler.Println("Write down an encrypted code, which will be decoded:")
				command, line = commandHandler.ReadCommand()

				if command == cli.Paste {
					line = string(clipboard.Read(clipboard.FmtText))
				}
			}

			commandHandler.Println(encryption.Decode(line))
		case cli.Paste:
			commandHandler.Printf("Pasted text is:\n%s\n", string(clipboard.Read(clipboard.FmtText)))
		case cli.Clear:
			commandHandler.Clear()
		default:
			commandHandler.Printf("(%s) Unknown command, try to use 'help'\n", strings.Split(line, " ")[0])
		}
	}

endOfMain:

	commandHandler.Printf("Bye, bye 🙋‍♂️\n")
}
