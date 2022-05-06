package install

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

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

	binariesToInstall, _ := assert.GetBinariesNotSatisfyingRequirements()
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
				ui.Abort("%s did not install successfully. Exiting...", binaryToInstall.Bin.Binary)
			}

			binPostInstall := binaryToInstall.Bin.PostInstall
			if binPostInstall != nil {
				ui.Message("Running post install command %s for bin: %s", binPostInstall.Command, binaryToInstall.Bin.Binary)
				_, err := util.ExecuteShellCommand(binPostInstall.Command, util.GetDefaultOptions())
				if err == nil {
					ui.Success("Post install %s successful", binaryToInstall.Bin.Binary)
				}
			}
		}
	}
	if dryRun {
		return
	}

	ui.Message("Now that installing is all done, making sure that you have installed all the requirements now...")
	binariesToInstall, _ = assert.GetBinariesNotSatisfyingRequirements()
	if len(binariesToInstall) > 0 {
		ui.Abort("Oh, no! I installed the binaries for you successfully however, it seems that the requirements are still not met. I am not sure how to fix this for you. Perhaps you need to improve your configuration file and try again?")
	}

	ui.Success("All binaries are installed and meet version requirements!")
}

func tryToInstallOrUpdateBinary(bin types.Bin, assertError assert.AssertError, isInteractive bool) bool {
	installersToChooseFrom := bin.Installers
	needsUpdated := assertError.NeedsUpdate

	if isInteractive {
		doNotCarePrompt := "No, I do not have a preference. Pick for me."

		prompt := promptui.Select{
			Label: fmt.Sprintf("Do you have a preference for what method is used to install %s?", bin.Binary),
			Items: append(installersToChooseFrom, doNotCarePrompt),
		}
		_, result, err := prompt.Run()
		ui.HandleError(err)

		if result == doNotCarePrompt {
			result = installersToChooseFrom[0] // choose one for you.
		}

		_, err = tryToInstallBinaryFromInstaller(bin, needsUpdated, result)
		ui.DebugError(err)

		didInstallSuccessfully := err == nil
		return didInstallSuccessfully
	} else {
		// TODO pick a preference of installing with the installer that is already installed. if you have gem but dont have brew, use gem.

		for _, command := range bin.Installers {
			_, err := tryToInstallBinaryFromInstaller(bin, needsUpdated, command)
			if err == nil {
				ui.Debug("Successfully installed %s", bin.Binary)
				return true
			}
		}

		return false
	}
}

func tryToInstallBinaryFromInstaller(bin types.Bin, needsUpdated bool, installerId string) (stdout string, err error) {
	ui.Debug("Trying to install %s with installer ID %s...", bin.Binary, installerId)

	installer := types.GetInstallerFromId(installerId)
	if installer == nil {
		err = fmt.Errorf("didn't find installer for installer id %s", installerId)
		return
	}

	if !util.IsBinInstalled(installer.Binary) {
		ui.Message("Before installing %s, you must first install the program that will install it for you. Installing %s for you...", bin.Binary, installer.Binary)

		stdout, err = tryToRunCommend(installer.InstallCommand)
		if err != nil {
			ui.DebugError(err)

			err = fmt.Errorf("error installing the installer %s", installer)
			return
		}
	}

	var installCommand string
	if needsUpdated {
		installCommand = getInstallCommandFromTemplate(bin, installer.UpdateTemplate)
	} else {
		installCommand = getInstallCommandFromTemplate(bin, installer.InstallTemplate)
	}

	stdout, err = tryToRunCommend(installCommand)
	if err == nil {
		ui.Success("%s installed %s successfully", ui.Emojis[":check_mark:"], bin.Binary)
	}

	return
}

func tryToRunCommend(command string) (stdout string, err error) {
	ui.Debug("running command: %s", command)

	shellCommandOptions := util.GetDefaultOptions()
	shellCommandOptions.StdoutToOS = true
	stdout, err = util.ExecuteShellCommand(command, shellCommandOptions)

	return
}

func getInstallCommandFromTemplate(bin types.Bin, installTemplate string) string {
	ui.Debug("Parsing install template for %s: %s", bin.Binary, installTemplate)
	installCommandTemplate, err := template.New("install command").Parse(installTemplate)
	ui.HandleError(err)

	var installCommandBuf bytes.Buffer
	installCommandTemplate.Execute(&installCommandBuf, bin)

	installCommand := installCommandBuf.String()
	ui.Debug("Command to install for %s is: %s", bin.Binary, installCommand)

	return installCommand
}
