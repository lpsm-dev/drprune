package commands

import (
	"github.com/lpmatos/drprune/internal/version"
	"github.com/spf13/cobra"
)

var (
	short  bool // Local flag - if true print just the version number of CLI.
	pretty bool // Local flag - if true show more details about the current version of CLI.
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Version outputs the version of CLI",
	Long:    `Version outputs the version of the drprune binary that is in use.`,
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			version.GetShortDetails()
		} else {
			version.ShowVersion(pretty)
		}
	},
}

func init() {
	versionCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Print just the version number of drprune CLI")
	versionCmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "Show more details about the current version of drprune CLI")
	rootCmd.AddCommand(versionCmd)
}
