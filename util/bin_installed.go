package util

import (
	"fmt"

	"github.com/levibostian/bins/ui"
)

func IsBinInstalled(bin string) bool {
	stdout, err := ExecuteShellCommand(fmt.Sprintf("command -v %s", bin))
	ui.Debug("Checking if %s binary installed. stdout %s, err %v", bin, stdout, err)

	isInstalled := err == nil

	ui.Debug("%s binary installed: %t", bin, isInstalled)

	return isInstalled
}
