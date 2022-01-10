package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// UpgradeCmd downloads and installs the most recent version of the CLI.
var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"u"},
	Short:   "Upgrade to the latest version of the CLI.",
	Long: `Upgrade to the latest version of the CLI.

This finds and downloads the latest release, if you don't
already have it.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Upgrade")
	},
}

func init() {
	rootCmd.AddCommand(UpgradeCmd)
}
