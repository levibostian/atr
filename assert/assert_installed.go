package assert

import (
	"github.com/levibostian/atr/types"
	"github.com/levibostian/atr/util"
)

func AssertBinariesInstalled(bins types.Bins) []AssertError {
	var assertErrors []AssertError

	for _, bin := range bins {
		if !util.IsBinInstalled(bin.Binary) {
			assertErrors = append(assertErrors, AssertError{
				Bin:         bin,
				IsInstalled: false,
			})
		}
	}

	return assertErrors
}
