package assert

import (
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
)

type AssertError struct {
	Bin              types.Bin
	IsInstalled      bool
	NeedsUpdate      bool
	InstalledVersion *string
	RequiredVersion  *string
}

func RunCommand() {
	assertErrors, validBins := GetBinariesNotSatisfyingRequirements()

	for _, assertError := range assertErrors {
		if !assertError.IsInstalled {
			ui.Error("%s %s is not installed", ui.Emojis[":red_x:"], assertError.Bin.Binary)
		}
		if assertError.NeedsUpdate {
			ui.Error("%s %s is installed with version %s but requires %s", ui.Emojis[":red_x:"], assertError.Bin.Binary, *assertError.InstalledVersion, *assertError.RequiredVersion)
		}
	}
	for _, validBin := range validBins {
		ui.Success("%s %s is installed and has a valid version", ui.Emojis[":check_mark:"], validBin.Binary)
	}

	if len(assertErrors) > 0 {
		ui.Abort("Fix issues listed above and try again.")
	}
	ui.Success("All binaries installed with required version!")
}

func GetBinariesNotSatisfyingRequirements() ([]AssertError, []types.Bin) {
	requiredBins := GetRequiredBins()

	return AssertBinariesInstalledAndVersionMet(requiredBins)
}
