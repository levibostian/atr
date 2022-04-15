package cmd

import (
	"github.com/levibostian/bins/assert"
	"github.com/spf13/cobra"
)

var assertCmd = &cobra.Command{
	Use:   "assert",
	Short: "Assert the local environment has all required binaries installed.",
	Long: `Assert the local environment has all required binaries (and required version) installed.
	
If the requirements are not met, tool will assist the user to install or upgrade their binaries to meet requirements`,
	Run: func(cmd *cobra.Command, args []string) {
		assert.AssertThenRun()
	},
}

func init() {
	rootCmd.AddCommand(assertCmd)
}
