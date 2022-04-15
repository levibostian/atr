package assert

import (
	"github.com/levibostian/atr/types"
	"github.com/levibostian/atr/ui"
	"github.com/spf13/viper"
)

func GetRequiredBins() types.Bins {
	var bins types.Bins
	viper.UnmarshalKey("bins", &bins)
	ui.Debug("bins from config file: %v", bins)

	return bins
}
