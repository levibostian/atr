package cmd

import (
	"github.com/levibostian/bins/assert"
	"github.com/levibostian/bins/ui"
	"github.com/spf13/cobra"
)

var assertCmd = &cobra.Command{
	Use:   "assert",
	Short: "Check if this computer has all required binaries installed and version requirements met.",
	Long: `Check if this computer has all required binaries installed and version requirements met.
	
This command will only exit successfully, or as a failure depending on if all binaries are installed or not.

This command is good to use in a git hook, for example, to make sure all tools are installed before trying to run them.`,
	Run: func(cmd *cobra.Command, args []string) {
		AssertCommandRun()
	},
}

func init() {
	rootCmd.AddCommand(assertCmd)
}

// TODO I am trying to make this tool accessible by program as well as via Go code. So, I am making public functions availbale to call for each command.
// I want to be able to disable all logging to stdout when not run via CLI. so, when any of the Cli commands run, I want to enable a flag to endable any logging. is there a cobra function that gets run for any of the cli functions?
func AssertCommandRun() {
	ui.Message("Welcome! I am going to check to see if you have the required programs (and required version of those programs) installed.")

	assert.RunCommand()
}
