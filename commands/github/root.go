package github

import (
	"github.com/spf13/cobra"
)

var token string
var name string
var container string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gh",
		Short: "Perform github operations",
		Long:  ``,
	}
	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitHub API Token")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "lpmatos", "GitHub User/Organization Name")
	rootCmd.PersistentFlags().StringVarP(&container, "container", "c", "", "GitHub Container Name")
	return rootCmd
}
