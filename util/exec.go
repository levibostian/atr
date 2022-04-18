package util

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

type ExecuteOptions struct {
	EnvVars    []string
	StdoutToOS bool
}

func GetDefaultOptions() ExecuteOptions {
	return ExecuteOptions{
		EnvVars:    []string{},
		StdoutToOS: false,
	}
}

func ExecuteShellCommand(command string, options ExecuteOptions) (stdout string, err error) {
	commandSplit := strings.Split(command, " ")

	cmd := exec.Command(commandSplit[0], commandSplit[1:]...)
	cmd.Env = os.Environ()
	for _, envVar := range options.EnvVars {
		cmd.Env = append(cmd.Env, envVar)
	}

	var outputBuffer bytes.Buffer
	multiWriter := io.MultiWriter(&outputBuffer)
	if options.StdoutToOS {
		multiWriter = io.MultiWriter(os.Stdout, &outputBuffer)
	}

	cmd.Stdout = multiWriter
	cmd.Stderr = multiWriter

	err = cmd.Run()
	stdout = StringTrimAll(outputBuffer.String())

	return
}
