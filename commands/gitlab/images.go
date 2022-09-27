package gitlab

import (
	"fmt"

	"github.com/ci-monk/drprune/internal/constants"
	"github.com/spf13/cobra"
)

var dryRun, untagged bool
var olderThan int

func NewCmdImages() *cobra.Command {
	var imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "Perform prune images operation on GitLab Registry (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(constants.ASCIIPrune)

			checkCmdParams()

			fmt.Println("Gl images")
		},
	}
	imagesCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Controlling whether to execute the action as a dry-run")
	imagesCmd.PersistentFlags().BoolVarP(&untagged, "untagged", "u", false, "Boolean controlling whether untagged versions should be pruned")
	imagesCmd.PersistentFlags().IntVarP(&olderThan, "older-than", "o", 0, "Minimum age in days of a version before it is pruned")
	return imagesCmd
}
