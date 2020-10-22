package loginshell

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func Shell() (string, error) {
	dir := "/Local/Default/Users/" + os.Getenv("USER")
	out, err := exec.Command("dscl", "localhost", "-read", dir, "UserShell").Output()
	if err != nil {
		return "", fmt.Errorf("dscl failed: %w", err)
	}
	re := regexp.MustCompile("UserShell: (/[^ ]+)\n")
	matched := re.FindStringSubmatch(string(out))
	shell := matched[1]
	if shell == "" {
		return "", fmt.Errorf("invalid output: %s", string(out))
	}
	return shell, nil
}
