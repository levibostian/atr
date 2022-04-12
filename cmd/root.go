package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/levibostian/atr/assert"
	"github.com/levibostian/atr/store"
	"github.com/levibostian/atr/ui"
)

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atr",
	Short: "Assert the local environment has all required binaries installed.",
	Long: `Assert the local environment has all required binaries (and required version) installed.
	
If the requirements are not met, tool will assist the user to install or upgrade their binaries to meet requirements`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		assert.AssertThenRun()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	ui.HandleError(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .atr.yml)")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug statements. Used for debugging program for bug reports and development. (default false)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // user defined it as an arg
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(".atr.yml")
	}
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	ui.HandleError(err)

	ui.Message("Using config file: " + viper.ConfigFileUsed())

	store.SetCliConfig(debug)
}
