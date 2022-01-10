package gitlab

import (
	"fmt"
	"net/url"
	"path"

	log "github.com/lpmatos/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	api *gitlab.Client
}

func NewClient(baseURL, token string, check bool) (*Client, error) {
	client := &Client{}

	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("can't parse URL: %v", err)
	}
	url.Path = path.Join(url.Path, "/api/v4")
	url.Scheme = "https"

	if baseURL != "" {
		client.api, err = gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	} else {
		client.api, err = gitlab.NewClient(token)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create gitlab client: %v", err)
	}

	if !check {
		version, _, err := client.api.Version.GetVersion()
		if err != nil {
			return nil, err
		}
		log.Debug(version.Version)
	}

	return client, nil
}

func (client *Client) GetUsername() (string, error) {
	user, _, err := client.api.Users.CurrentUser()
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (client *Client) GetGroupAllRegistryRepositories() {
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

func (client *Client) GetProjectAllRegistryRepositories() {
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
