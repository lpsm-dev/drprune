package gitlab

import (
	"fmt"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

// GetGroupAllRegistryRepositories obtém todos os repositórios do registry do grupo.
func (client *GitLabClient) GetGroupAllRegistryRepositories(groupPath string) {
	page := 0
	perPage := 20

	for {
		groupRepos, resp, err := client.api.ContainerRegistry.ListGroupRegistryRepositories(
			groupPath,
			&gitlab.ListRegistryRepositoriesOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			},
		)
		if err != nil {
			log.Fatalf("Erro ao obter repositórios de registry do grupo: %v", err)
		}

		for _, value := range groupRepos {
			fmt.Printf("> Localização: %v\n", value.Location)
		}
		fmt.Println("==============================")

		page++
		if resp.TotalPages == page || len(groupRepos) == 0 {
			break
		}
	}
}
