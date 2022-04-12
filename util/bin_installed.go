package util

import (
	"os/exec"

	"github.com/levibostian/atr/ui"
)

func IsBinInstalled(bin string) bool {
	cmd := exec.Command("command", "-v", bin)
	stdout, err := cmd.Output()
	ui.Debug("Checking if %s binary installed. stdout %s, err %v", bin, stdout, err)

	isInstalled := err == nil

	ui.Debug("%s binary installed: %t", bin, isInstalled)

	return isInstalled
}
