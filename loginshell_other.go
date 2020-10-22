//+build !linux,!openbsd,!freebsd,!darwin,!plan9,!windows

package loginshell

import (
	"errors"
	"os"
)

func Shell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", errors.New("SHELL not defined")
	}
	return shell, nil
}
