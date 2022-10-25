package gitlab

import (
	"fmt"
	"net/url"
	"path"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

func NewClient(baseURL, token string, check bool) (*GitLabClient, error) {
	// Getting the GitLab client.
	client := &GitLabClient{}

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

func (client *GitLabClient) GetUsername() (string, error) {
	user, _, err := client.api.Users.CurrentUser()
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
