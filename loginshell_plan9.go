package loginshell

import (
	"fmt"
	"os"
)

func Shell() (string, error) {
	if _, err := os.Stat("/dev/osversion"); err != nil {
		return "", fmt.Errorf("/dev/osversion check failed: %w", err)
	}
	return "/bin/rc", nil
}
