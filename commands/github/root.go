package github

import (
	"os"

	"github.com/spf13/cobra"
)

var token, username, container string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gh",
		Short: "Perform github operations",
		Long:  ``,
	}

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", os.Getenv("GH_TOKEN"), "GitHub API Token (*)")
	rootCmd.PersistentFlags().StringVarP(&username, "name", "n", os.Getenv("GH_USERNAME"), "GitHub User/Organization Name (*)")
	rootCmd.PersistentFlags().StringVarP(&container, "container", "c", os.Getenv("GH_CONTAINER"), "GitHub Container Name (*)")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	return rootCmd
}
