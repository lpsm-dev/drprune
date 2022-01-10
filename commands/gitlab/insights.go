package gitlab

import (
	log "github.com/lpmatos/drprune/internal/log"
	gl "github.com/lpmatos/drprune/pkg/gitlab"
	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitLab Registry (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gl.NewClient(url, token, false)
			if err != nil {
				log.Fatal(err)
			}

			user, err := client.GetUsername()
			if err != nil {
				log.Fatal(err)
			}
			log.Debug(user)

			client.GetGroupRegistry()

		},
	}
	return insightsCmd
}
