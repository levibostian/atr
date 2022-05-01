package types

import (
	"github.com/levibostian/bins/ui"
	fp "github.com/repeale/fp-go"
	"github.com/spf13/viper"
)

type Installer struct {
	Id              string
	Binary          string
	InstallCommand  string
	InstallTemplate string
	UpdateTemplate  string
}

var Installers []Installer = GetInstallersFromConfig()

func GetInstallersFromConfig() []Installer {
	var installers []Installer
	viper.UnmarshalKey("installers", &installers)
	ui.Debug("installers from config file: %v", installers)

	return installers
}

func GetInstallerFromId(installerId string) *Installer {
	foundInstaller := fp.Filter(func(installer Installer) bool { return installer.Id == installerId })(Installers)
	if len(foundInstaller) <= 0 {
		return nil
	}

	return &foundInstaller[0]
}
