package gitlab

import (
	"fmt"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

func (client *GitLabClient) GetProjectAllRegistryRepositories() {
	page := 0
	for {
		projectRepos, resp, err := client.api.ContainerRegistry.ListProjectRegistryRepositories(
			"surfe/360cel/api/chip",
			&gitlab.ListRegistryRepositoriesOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: 20,
				},
			})
		if err != nil {
			log.Fatal(err)
		}

		for _, value := range projectRepos {
			fmt.Printf("> Location: %v\n", value.Location)
		}
		fmt.Println("==============================")

		page += 1
		if resp.TotalPages == page || len(projectRepos) == 0 {
			break
		}
	}
}
