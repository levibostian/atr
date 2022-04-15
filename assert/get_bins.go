package assert

import (
	"github.com/levibostian/bins/types"
	"github.com/levibostian/bins/ui"
	"github.com/spf13/viper"
)

func GetRequiredBins() types.Bins {
	var bins types.Bins
	viper.UnmarshalKey("bins", &bins)
	ui.Debug("bins from config file: %v", bins)

	return bins
}
