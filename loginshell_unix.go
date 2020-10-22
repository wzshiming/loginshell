//+build linux,openbsd,freebsd

package loginshell

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"
)

func Shell() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	out, err := exec.Command("getent", "passwd", user.Uid).Output()
	if err != nil {
		return "", fmt.Errorf("getent failed: %w", err)
	}
	ent := strings.Split(strings.TrimSuffix(string(out), "\n"), ":")
	return ent[6], nil
}
