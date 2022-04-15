package assert

import (
	"fmt"

	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
)

type AssertError struct {
	Bin              types.Bin
	IsInstalled      bool
	InstalledVersion *string
	RequiredVersion  *string
}

func AssertThenRun() {
	assertErrors := GetBinariesNotSatisfyingRequirements()

	if len(assertErrors) > 0 {
		ui.Message(getErrorMessageFromAssertErrors(assertErrors))

		ui.Abort("Fix issues listed above and try again.")
	}

	ui.Success("All binaries installed with required version!")
}

func GetBinariesNotSatisfyingRequirements() []AssertError {
	requiredBins := GetRequiredBins()
	binariesNotInstalled := AssertBinariesInstalled(requiredBins)
	binaryVersionRequirementsNotMet := AssertBinariesVersionMet(requiredBins)

	return append(binariesNotInstalled, binaryVersionRequirementsNotMet...)
}

func getErrorMessageFromAssertErrors(assertErrors []AssertError) string {
	var errorMessage string = ""

	for _, assertError := range assertErrors {
		if !assertError.IsInstalled {
			errorMessage += fmt.Sprintf("%s is not installed\n", assertError.Bin.Binary)
		}

		if assertError.IsInstalled && assertError.RequiredVersion != nil {
			errorMessage += fmt.Sprintf("%s version is %s, but is required to be: %s\n", assertError.Bin.Binary, *assertError.InstalledVersion, *assertError.RequiredVersion)
		}
	}

	return errorMessage
}
