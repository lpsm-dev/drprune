package gitlab

import (
	"fmt"
	"net/url"
	"path"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/xanzy/go-gitlab"
)

// GitLabClient representa um cliente para o GitLab.
type GitLabClient struct {
	api *gitlab.Client
}

// GetApiClient retorna um ponteiro para o objeto gitlab.Client do cliente GitLab.
func (client *GitLabClient) GetApiClient() *gitlab.Client {
	// Retorna o campo api da struct GitLabClient.
	return client.api
}

// NewClient cria um novo cliente para o GitLab.
// host é a URL base do GitLab, token é um token de acesso para autenticação e check é um indicador de se a verificação de versão é necessária.
func NewClient(host, token string, check bool) (*GitLabClient, error) {
	client := &GitLabClient{}

	// Analisa a URL base do GitLab.
	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("não é possível analisar a url: %v", err)
	}
	u.Path = path.Join(u.Path, "/api/v4")
	u.Scheme = "https"

	// Cria um novo cliente GitLab com a URL base e o token de acesso.
	if host != "" {
		client.api, err = gitlab.NewClient(token, gitlab.WithBaseURL(host))
	} else {
		client.api, err = gitlab.NewClient(token)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao criar o cliente do gitlab: %v", err)
	}
	log.Debugln("cliente do gitlab inicializado!")

	// Verifica a versão do cliente do GitLab, se necessário.
	if check {
		version, _, err := client.api.Version.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("falha ao pegar versão do gitlab: %v", err)
		}
		log.Debugf("estamos na versão %s do gitlab", version.Version)
	}

	return client, nil
}
