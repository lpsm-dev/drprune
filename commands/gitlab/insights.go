package gitlab

import (
	"fmt"

	"github.com/ci-monk/drprune/internal/constants"
	log "github.com/ci-monk/drprune/internal/log"
	gl "github.com/ci-monk/drprune/pkg/gitlab"
	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitLab Registry (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(constants.ASCIInsights)

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
