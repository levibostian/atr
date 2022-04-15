package install

import (
	"bytes"
	"text/template"

	"github.com/levibostian/atr/assert"
	"github.com/levibostian/atr/types"
	"github.com/levibostian/atr/ui"
	"github.com/levibostian/atr/util"
)

func RunCommand(dryRun bool) {
	ui.Debug("Running install command. dry-run %v", dryRun)

	binariesToInstall := assert.GetBinariesNotSatisfyingRequirements()
	if len(binariesToInstall) <= 0 {
		ui.Success("All binaries are installed and meet version requirements. Nothing to install!")
		return
	}

	if dryRun {
		ui.Message("Command run in dry run mode. Printing binaries that would be installed:")
	}
	for _, binaryToInstall := range binariesToInstall {
		if dryRun {
			ui.Message("• %s", binaryToInstall.Bin.Binary)
		} else {
			didInstallSuccessfully := tryToInstallBinary(binaryToInstall.Bin)
			if !didInstallSuccessfully {
				ui.Abort("%s did not install successfully. Exiting...")
			}
		}
	}
	if dryRun {
		return
	}

	binariesToInstall = assert.GetBinariesNotSatisfyingRequirements()
	if len(binariesToInstall) > 0 {
		ui.Abort("Oh, no! I installed the binaries for you successfully however, it seems that the requirements are still not met. I am not sure how to fix this for you. Perhaps you need to improve your configuration file and try again?")
	}

	ui.Success("All binaries are installed and meet version requirements!")
}

func tryToInstallBinary(bin types.Bin) bool {
	for _, installTemplate := range bin.InstallCommand {
		_, err := tryToInstallBinaryFromTemplate(bin, installTemplate)
		if err == nil {
			ui.Debug("Successfully installed %s", bin.Binary)
			return true
		}
	}

	return false
}

func tryToInstallBinaryFromTemplate(bin types.Bin, installTemplate string) (stdout string, err error) {
	ui.Debug("Parsing install template for %s: %s", bin.Binary, installTemplate)
	installCommandTemplate, err := template.New("install command").Parse(installTemplate)
	ui.HandleError(err)

	var installCommandBuf bytes.Buffer
	installCommandTemplate.Execute(&installCommandBuf, struct {
		Binary  string
		Version string
	}{bin.Binary, bin.GetVersion()})

	installCommand := installCommandBuf.String()
	ui.Debug("Command to install %s: %s", bin.Binary, installCommand)

	progressBar := ui.MessageProgress("Installing %s with command %s", bin.Binary, installCommand)
	stdout, err = util.ExecuteShellCommand(installCommand)
	progressBar.Done()
	if err == nil {
		ui.Success("%s Installed %s successfully", ui.Emojis[":check_mark:"], bin.Binary)
	}

	return
}
