package gitlab

import (
	"fmt"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

func (client *GitLabClient) GetGroupAllRegistryRepositories() {
	page := 0
	for {
		groupRepos, resp, err := client.api.ContainerRegistry.ListGroupRegistryRepositories(
			"surfe",
			&gitlab.ListRegistryRepositoriesOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: 20,
				},
			})
		if err != nil {
			log.Fatal(err)
		}

		for _, value := range groupRepos {
			fmt.Printf("> Location: %v\n", value.Location)
		}
		fmt.Println("==============================")

		page += 1
		if resp.TotalPages == page || len(groupRepos) == 0 {
			break
		}
	}
}
