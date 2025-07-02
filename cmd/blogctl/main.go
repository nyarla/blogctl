package main

import (
	"os"
)

const version = "0.0.1"
const help = `blogctl - the cli client for hatenablog.

Commands:

init - init .env file.
help - show blogctl help.
version - show blogctl version.
`

func main() {
	if len(os.Args) > 1 {
		var action = os.Args[1]

		switch action {
		case "init":
			doInitEnv()
		case "version", "-v", "--version":
			showVersion()
		case "help", "-h", "--help":
			showHelp()
		default:
			showUnsupportedCommand()
		}
	}

	showHelp()
}
