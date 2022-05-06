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

var bundledInstallers = []Installer{
	{
		Id:              "brew",
		Binary:          "brew",
		InstallCommand:  "/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"",
		InstallTemplate: "brew install {{.Binary}}",
		UpdateTemplate:  "brew upgrade {{.Binary}}",
	}, {
		Id:     "gem",
		Binary: "gem",
		// Do not try to install the ruby programming language for you. If it isn't installed, try another method.
		InstallCommand:  "false",
		InstallTemplate: "gem install {{.Binary}}",
		UpdateTemplate:  "gem update {{.Binary}}",
	},
}

func GetInstallersFromConfig() []Installer {
	var configFileInstallers []Installer
	viper.UnmarshalKey("installers", &configFileInstallers)
	ui.Debug("installers from config file: %v", configFileInstallers)

	var combinedInstallers = bundledInstallers
	// combine the bundled + config file installers.
	// allow user to override bundled.
	for _, configFileInstaller := range configFileInstallers {
		for index, bundledInstaller := range combinedInstallers {
			if configFileInstaller.Binary == bundledInstaller.Binary {
				combinedInstallers[index] = configFileInstaller
			} else {
				combinedInstallers = append(combinedInstallers, configFileInstaller)
			}
		}
	}

	return combinedInstallers
}

func GetInstallerFromId(installerId string) *Installer {
	foundInstaller := fp.Filter(func(installer Installer) bool { return installer.Id == installerId })(Installers)
	if len(foundInstaller) <= 0 {
		return nil
	}

	return &foundInstaller[0]
}
