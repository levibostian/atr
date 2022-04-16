package assert

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
	"github.com/levibostian/bins/util"
)

func AssertBinariesVersionMet(bins types.Bins) []AssertError {
	var assertErrors []AssertError

	for _, bin := range bins {
		installedVersion := getBinInstalledVersion(bin)
		if installedVersion == nil {
			ui.Debug("Could not determine version of %s", bin.Binary)
			// TODO this needs fixed. as of now, if it's installed but cant find version, we will not report it as a problem to fix in assert. this needs changed.

			// we will assume the bin is not installed so just ignore it
			continue
		}
		installedVersionString := installedVersion.String()

		requiredVersionConstraint, err := semver.NewConstraint(bin.Version.Requirement)
		if err != nil {
			ui.DebugError(err)
			continue
		}

		ui.Debug("Checking if %s meets required version %s with installed: %s", bin.Binary, bin.Version.Requirement, installedVersionString)
		isBinaryVersionRequirementMet := requiredVersionConstraint.Check(installedVersion)
		if !isBinaryVersionRequirementMet {
			ui.Debug("%s does not meet version requirement", bin.Binary)

			assertErrors = append(assertErrors, AssertError{
				Bin:              bin,
				IsInstalled:      true,
				InstalledVersion: &installedVersionString,
				RequiredVersion:  &bin.Version.Requirement,
			})
			continue
		}

		ui.Debug("%s does meet version requirement", bin.Binary)
	}

	return assertErrors
}

func getBinInstalledVersion(bin types.Bin) *semver.Version {
	var getVersionCommand string = fmt.Sprintf("%s --version", bin.Binary)
	if bin.Version.Command != nil {
		getVersionCommand = *bin.Version.Command
	}

	stdout, err := util.ExecuteShellCommand(getVersionCommand, bin.Version.EnvVars)
	ui.Debug("Getting version of %s. from stdout %s, err %v", bin.Binary, stdout, err)
	if err != nil {
		return nil
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
		return nil
	}

	ui.Debug("Version parsed for %s: %s", bin.Binary, versionResult.String())

	return versionResult
}
