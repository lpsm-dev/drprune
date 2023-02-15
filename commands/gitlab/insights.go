package gitlab

import (
	"fmt"

	"github.com/ci-monk/drprune/internal/consts"
	log "github.com/ci-monk/drprune/internal/log"
	gl "github.com/ci-monk/drprune/pkg/gitlab"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitLab Registry (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(consts.ASCIInsights)

			checkCmdParams()

			gitlabGroup, gitlabProject := "surfe", "surfe/360cel/api/chip"

			client, err := gl.NewClient(url, token, false)
			if err != nil {
				log.Fatal(err)
			}

			// ===========================================================================

			user, err := client.GetUsername()
			if err != nil {
				log.Fatal(err)
			}
			log.Debug(user)

			// ===========================================================================

			pterm.DefaultSection.Println("GitLab get all group registry repositories")
			groupRegistryRepos, err := client.GetGroupAllRegistryRepositories(gitlabGroup)
			if err != nil {
				log.Errorf("list all project registry repositories: %v", err)
			}
			// Loop each registry repo of group
			for _, registryRepo := range groupRegistryRepos {
				fmt.Printf("> Registry Repo Location: %s\n", registryRepo.Location)
			}

			fmt.Println()

			// ===========================================================================

			pterm.DefaultSection.Println("GitLab get all project registry repositories")
			projectRegistryRepos, err := client.GetProjectAllRegistryRepositories(gitlabProject)
			if err != nil {
				log.Errorf("list all project registry repositories: %v", err)
			}
			// Loop each registry repo of project
			for _, registryRepo := range projectRegistryRepos {
				fmt.Printf("> Registry Repo Location: %s\n", registryRepo.Location)
			}
		},
	}
	return insightsCmd
}
