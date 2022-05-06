package types

import (
	"github.com/levibostian/bins/ui"
	"github.com/spf13/viper"
)

type Bin struct {
	Binary  string
	Version struct {
		Requirement    string
		Command        *string
		CommandEnvVars *[]string
	}
	Installers  []string
	PostInstall *struct {
		Command string
	}
}

type Bins = []Bin

func GetBinsFromConfig() Bins {
	var bins Bins
	viper.UnmarshalKey("bins", &bins)
	ui.Debug("bins from config file: %v", bins)

	return bins
}
