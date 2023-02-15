package gitlab

import (
	"fmt"

	"github.com/ci-monk/drprune/internal/consts"
	log "github.com/ci-monk/drprune/internal/log"
	gl "github.com/ci-monk/drprune/pkg/gitlab"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// NewCmdInsights cria um novo comando Cobra para executar os insights do GitLab.
func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitLab Registry (registry.gitlab.com)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Imprime o banner ASCI
			fmt.Printf(consts.ASCIInsights)

			// Verifica os parâmetros do comando
			checkCmdParams()

			// Define o grupo e o projeto do GitLab
			gitlabGroup, gitlabProject := "surfe", "surfe/360cel/api/chip"

			// Cria um novo cliente GitLab
			client, err := gl.NewClient(url, token, false)
			if err != nil {
				log.Fatal(err)
			}

			// Obtém todos os registros de repositórios do grupo GitLab e imprime no console
			pterm.DefaultSection.Println("GitLab get all group registry repositories")
			groupRegistryRepos, err := client.GetGroupAllRegistryRepositories(gitlabGroup)
			if err != nil {
				log.Errorf("list all group registry repositories: %v", err)
			}

			for _, registryRepo := range groupRegistryRepos {
				fmt.Printf("> Registry Repo Location: %s\n", registryRepo.Location)
			}

			fmt.Println()

			// Obtém todos os registros de repositórios do projeto GitLab e imprime no console
			pterm.DefaultSection.Println("GitLab get all project registry repositories")
			projectRegistryRepos, err := client.GetProjectAllRegistryRepositories(gitlabProject)
			if err != nil {
				log.Errorf("list all project registry repositories: %v", err)
			}
			for _, registryRepo := range projectRegistryRepos {
				fmt.Printf("> Registry Repo Location: %s\n", registryRepo.Location)
			}
		},
	}
	return insightsCmd
}
