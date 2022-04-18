package assert

import (
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/util"
)

func AssertBinariesInstalledAndVersionMet(bins types.Bins) (errors []AssertError, validBins []types.Bin) {
	var binariesInstalled []types.Bin
	for _, bin := range bins {
		if !util.IsBinInstalled(bin.Binary) {
			errors = append(errors, AssertError{
				Bin:         bin,
				IsInstalled: false,
			})
		} else {
			binariesInstalled = append(binariesInstalled, bin)
		}
	}

	binariesVersionNotMet, validBins := AssertBinariesVersionMet(binariesInstalled)
	errors = append(errors, binariesVersionNotMet...)

	return
}
