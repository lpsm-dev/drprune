package commands

import "github.com/spf13/cobra"

var organizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"org"},
	Short:   "Prune organization images",
	Long:    ``,
}

func init() {
	rootCmd.AddCommand(organizationCmd)
}
