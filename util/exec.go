package util

import (
	"os"
	"os/exec"
	"strings"
)

func ExecuteShellCommand(command string, envVars *[]string) (stdout string, err error) {
	commandSplit := strings.Split(command, " ")

	cmd := exec.Command(commandSplit[0], commandSplit[1:]...)
	cmd.Env = os.Environ()
	if envVars != nil {
		for _, envVar := range *envVars {
			cmd.Env = append(cmd.Env, envVar)
		}
	}

	stdoutBytes, err := cmd.Output()
	stdout = StringTrimAll(string(stdoutBytes))

	return
}
