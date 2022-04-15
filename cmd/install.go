package cmd

import (
	"github.com/levibostian/atr/install"
	"github.com/spf13/cobra"
)

var dryRunInstall = false

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "No interaction installing of all dependencies.",
	Long:  `No interaction installing of all dependencies. Great for running on a CI server to quickly and easily install all of your dependencies needed. Not recommended for development machine as the interactive method is preferred there.`,
	Run: func(cmd *cobra.Command, args []string) {
		install.RunCommand(dryRunInstall)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVar(&dryRunInstall, "dry-run", false, "display actions to be done, but do not actually install")
}
