package gitlab

import (
	"fmt"
	"net/url"
	"path"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

// GitLabClient é um cliente para o GitLab.
type GitLabClient struct {
	api *gitlab.Client
}

// NewClient cria um novo cliente para o GitLab.
func NewClient(host, token string, check bool) (*GitLabClient, error) {
	client := &GitLabClient{}

	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("não é possível analisar o URL: %v", err)
	}
	u.Path = path.Join(u.Path, "/api/v4")
	u.Scheme = "https"

	if host != "" {
		client.api, err = gitlab.NewClient(token, gitlab.WithBaseURL(host))
	} else {
		client.api, err = gitlab.NewClient(token)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao criar o cliente do GitLab: %v", err)
	}

	if !check {
		version, _, err := client.api.Version.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("falha ao obter a versão do GitLab: %v", err)
		}
		log.Debug(version.Version)
	}

	return client, nil
}

// GetUsername retorna o nome de usuário do usuário atual.
func (client *GitLabClient) GetUsername() (string, error) {
	user, _, err := client.api.Users.CurrentUser()
	if err != nil {
		return "", fmt.Errorf("falha ao obter o usuário atual: %v", err)
	}

	return user.Username, nil
}
