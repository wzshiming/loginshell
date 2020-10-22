package loginshell

import (
	"os"
)

func Shell() (string, error) {
	consoleApp := os.Getenv("COMSPEC")
	if consoleApp == "" {
		consoleApp = "cmd.exe"
	}
	return consoleApp, nil
}
