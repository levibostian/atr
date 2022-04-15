package util

import (
	"os/exec"
	"strings"
)

func ExecuteShellCommand(command string) (stdout string, err error) {
	commandSplit := strings.Split(command, " ")

	cmd := exec.Command(commandSplit[0], commandSplit[1:]...)
	stdoutBytes, err := cmd.Output()
	stdout = StringTrimAll(string(stdoutBytes))

	return
}
