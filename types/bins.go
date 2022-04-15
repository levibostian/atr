package types

import (
	"strings"
)

type Bin struct {
	Binary         string
	Version        string
	InstallCommand []string
}

type Bins = []Bin

func (bin Bin) GetVersion() string {
	return strings.NewReplacer("^", "", "~", "").Replace(bin.Version)
}
