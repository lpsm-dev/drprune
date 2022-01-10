package gitlab

import (
	"fmt"

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

			log.Infoln("List groups")

			client.GetGroupAllRegistryRepositories()

			fmt.Println()

			log.Infoln("List project")

			client.GetProjectAllRegistryRepositories()

		},
	}
	return insightsCmd
}
