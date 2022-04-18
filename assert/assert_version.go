package assert

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
	"github.com/levibostian/bins/util"
)

func AssertBinariesVersionMet(bins types.Bins) (errors []AssertError, validBins []types.Bin) {
	for _, bin := range bins {
		installedVersion := getBinInstalledVersion(bin)
		installedVersionString := installedVersion.String()

		requiredVersionConstraint, err := semver.NewConstraint(bin.Version.Requirement)
		if err != nil {
			ui.DebugError(err)
			continue
		}

		ui.Debug("Checking if %s meets required version %s with installed: %s", bin.Binary, bin.Version.Requirement, installedVersionString)
		isBinaryVersionRequirementMet := requiredVersionConstraint.Check(&installedVersion)
		if !isBinaryVersionRequirementMet {
			ui.Debug("%s does not meet version requirement", bin.Binary)

			errors = append(errors, AssertError{
				Bin:              bin,
				IsInstalled:      true,
				NeedsUpdate:      true,
				InstalledVersion: &installedVersionString,
				RequiredVersion:  &bin.Version.Requirement,
			})
			continue
		} else {
			validBins = append(validBins, bin)
		}

		ui.Debug("%s does meet version requirement", bin.Binary)
	}

	return
}

func getBinInstalledVersion(bin types.Bin) semver.Version {
	var getVersionCommand string = fmt.Sprintf("%s --version", bin.Binary)
	if bin.Version.Command != nil {
		getVersionCommand = *bin.Version.Command
	}

	shellCommandOptions := util.GetDefaultOptions()
	if bin.Version.EnvVars != nil {
		shellCommandOptions.EnvVars = *bin.Version.EnvVars
	}
	stdout, err := util.ExecuteShellCommand(getVersionCommand, shellCommandOptions)
	ui.Debug("Getting version of %s. from stdout %s, err %v", bin.Binary, stdout, err)
	if err != nil {
		ui.Abort("Could not determine version of %s. Check your configuration to help successfully get the version string. Here is stdout from trying to get version: %s", bin.Binary, stdout)
	}

	var stdoutWords []string
	for _, line := range strings.Split(string(stdout), "\n") {
		for _, word := range strings.Split(line, " ") {
			stdoutWords = append(stdoutWords, word)
		}
	}

	var versionResult *semver.Version
	for _, word := range stdoutWords {
		version, _ := semver.NewVersion(word)
		if version != nil {
			versionResult = version
			break
		}
	}

	if versionResult == nil {
		ui.Abort("Could not determine version of %s. Check your configuration to help successfully get the version string. Here is stdout from trying to get version: %s", bin.Binary, stdout)
	}

	ui.Debug("Version parsed for %s: %s", bin.Binary, versionResult.String())

	return *versionResult
}
