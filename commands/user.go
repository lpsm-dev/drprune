package commands

import "github.com/spf13/cobra"

var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Prune user images",
	Long:    ``,
}

func init() {
	rootCmd.AddCommand(userCmd)
}
