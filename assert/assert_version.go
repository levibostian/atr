package assert

import (
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/levibostian/atr/ui"
)

func AssertBinariesVersionMet(bins Bins) []AssertError {
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
			continue
		}

		ui.Debug("Checking if %s meets required version %s with installed: %s", bin.Binary, bin.Version, installedVersionString)
		isBinaryVersionRequirementMet := requiredVersionConstraint.Check(installedVersion)
		if !isBinaryVersionRequirementMet {
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
	cmd := exec.Command(bin, "--version")
	stdout, err := cmd.Output()
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
