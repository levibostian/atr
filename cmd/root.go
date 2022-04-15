package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/levibostian/bins/assert"
	"github.com/levibostian/bins/store"
	"github.com/levibostian/bins/ui"
)

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bins",
	Short: "Assert everyone on your team (or CI server) has binaries installed for project.",
	Long:  `Assert everyone on your team (or CI server) has binaries installed for project. Makes onboarding with a new project a more positive experience.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// default command is assert
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .bins.yml)")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug statements. Used for debugging program for bug reports and development. (default false)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // user defined it as an arg
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(".bins.yml")
	}
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	ui.HandleError(err)

	ui.Message("Using config file: " + viper.ConfigFileUsed())

	store.SetCliConfig(debug)
}
