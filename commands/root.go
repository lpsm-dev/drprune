package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var token string

var rootCmd = &cobra.Command{
	Use:                "ghcr-prune",
	Short:              "Prune old images on GitHub Container Registry (ghcr.io)",
	Long:               ``,
	DisableSuggestions: false,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(token)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitHub API Token")
}
