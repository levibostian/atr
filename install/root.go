package install

import (
	"os"

	"github.com/levibostian/bins/assert"
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
	"github.com/levibostian/bins/util"
	"github.com/manifoldco/promptui"
)

func RunCommand(dryRun bool) {
	isOnCI := os.Getenv("CI") != ""
	isInteractive := !isOnCI

	ui.Debug("Running install. interactive %v, dry-run %v", isInteractive, dryRun)

	binariesToInstall := assert.GetBinariesNotSatisfyingRequirements()
	if len(binariesToInstall) <= 0 {
		ui.Success("All binaries are installed and meet version requirements. Nothing to install!")
		return
	}

	if dryRun {
		ui.Message("Command run in dry run mode. Printing binaries that would be installed or updated:")
	}
	for _, binaryToInstall := range binariesToInstall {
		if dryRun {
			var suffix string = "(install)"
			if binaryToInstall.NeedsUpdate {
				suffix = "(update)"
			}

			ui.Message("• %s %s", binaryToInstall.Bin.Binary, suffix)
		} else {
			didInstallSuccessfully := tryToInstallOrUpdateBinary(binaryToInstall.Bin, binaryToInstall, isInteractive)
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

func tryToInstallOrUpdateBinary(bin types.Bin, assertError assert.AssertError, isInteractive bool) bool {
	var commandsToChooseFrom []string // depending on if we need to install or upgrade, get the options
	updateBin := assertError.NeedsUpdate
	for _, installMethod := range bin.InstallMethods {
		if updateBin && installMethod.UpdateCommand != nil {
			commandsToChooseFrom = append(commandsToChooseFrom, *installMethod.UpdateCommand)
		} else {
			commandsToChooseFrom = append(commandsToChooseFrom, installMethod.Command)
		}
	}

	if isInteractive && len(bin.InstallMethods) > 1 {
		prompt := promptui.Select{
			Label: "What method of installing do you prefer?",
			Items: commandsToChooseFrom,
		}
		_, result, err := prompt.Run()
		ui.HandleError(err)

		_, err = tryToInstallBinaryFromCommand(bin, result)

		didInstallSuccessfully := err == nil
		return didInstallSuccessfully
	} else {
		for _, command := range commandsToChooseFrom {
			_, err := tryToInstallBinaryFromCommand(bin, command)
			if err == nil {
				ui.Debug("Successfully installed %s", bin.Binary)
				return true
			}
		}

		return false
	}
}

func tryToInstallBinaryFromCommand(bin types.Bin, installCommand string) (stdout string, err error) {
	progressBar := ui.MessageProgress("Installing %s with command %s", bin.Binary, installCommand)
	stdout, err = util.ExecuteShellCommand(installCommand, nil)
	progressBar.Done()
	if err == nil {
		ui.Success("%s Installed %s successfully", ui.Emojis[":check_mark:"], bin.Binary)
	}

	return
}
