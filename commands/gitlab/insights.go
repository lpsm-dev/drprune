package gitlab

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights off your GitLab (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gl images")
		},
	}
	return insightsCmd
}
