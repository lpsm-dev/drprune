package github

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights off your GitHub (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gh insights")
		},
	}
	return insightsCmd
}
