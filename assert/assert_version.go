package assert

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/levibostian/atr/types"
	"github.com/levibostian/atr/ui"
	"github.com/levibostian/atr/util"
)

func AssertBinariesVersionMet(bins types.Bins) []AssertError {
	var assertErrors []AssertError

	for _, bin := range bins {
		installedVersion, err := getBinInstalledVersion(bin.Binary)
		if err != nil {
			// we will assume the bin is not installed so just ignore it
			continue
		}
		installedVersionString := installedVersion.String()

		requiredVersionConstraint, err := semver.NewConstraint(bin.Version)
		if err != nil {
			ui.DebugError(err)
			continue
		}

		ui.Debug("Checking if %s meets required version %s with installed: %s", bin.Binary, bin.Version, installedVersionString)
		isBinaryVersionRequirementMet := requiredVersionConstraint.Check(installedVersion)
		if !isBinaryVersionRequirementMet {
			ui.Debug("%s does not meet version requirement", bin.Binary)

			assertErrors = append(assertErrors, AssertError{
				Bin:              bin,
				IsInstalled:      true,
				InstalledVersion: &installedVersionString,
				RequiredVersion:  &bin.Version,
			})
			continue
		}

		ui.Debug("%s does meet version requirement", bin.Binary)
	}

	return assertErrors
}

func getBinInstalledVersion(bin string) (*semver.Version, error) {
	stdout, err := util.ExecuteShellCommand(fmt.Sprintf("%s --version", bin))
	ui.Debug("Getting version of %s. stdout %s, err %v", bin, stdout, err)
	if err != nil {
		return nil, err
	}

	version, err := semver.NewVersion(strings.TrimSpace(string(stdout)))
	if err != nil {
		ui.DebugError(err)
		return nil, err
	}
	ui.Debug("Version parsed for %s: %s", bin, version.String())

	return version, nil
}
