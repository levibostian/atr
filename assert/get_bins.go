package assert

import (
	"github.com/levibostian/atr/ui"
	"github.com/spf13/viper"
)

func GetRequiredBins() Bins {
	var bins Bins
	viper.UnmarshalKey("bins", &bins)
	ui.Debug("bins from config file %v", bins)

	return bins
}
