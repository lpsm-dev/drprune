package gitlab

import (
	"fmt"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

// GetProjectAllRegistryRepositories obtém todos os repositórios do registry do projeto.
func (client *GitLabClient) GetProjectAllRegistryRepositories(projectPath string) {
	page := 0
	perPage := 20

	for {
		projectRepos, resp, err := client.api.ContainerRegistry.ListProjectRegistryRepositories(
			projectPath,
			&gitlab.ListRegistryRepositoriesOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			},
		)
		if err != nil {
			log.Fatalf("Erro ao obter repositórios de registry do projeto: %v", err)
		}

		for _, value := range projectRepos {
			fmt.Printf("> Localização: %v\n", value.Location)
		}
		fmt.Println("==============================")

		page++
		if resp.TotalPages == 0 || resp.TotalPages == page || len(projectRepos) == 0 {
			break
		}
	}
}