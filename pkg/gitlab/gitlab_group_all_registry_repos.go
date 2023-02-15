package gitlab

import (
	"fmt"
	"net/http"

	"github.com/xanzy/go-gitlab"
)

// GetGroupAllRegistryRepositories obtém uma lista de todos os repositórios do registro de contêineres do GitLab para um determinado grupo.
func (client *GitLabClient) GetGroupAllRegistryRepositories(groupPath string) ([]*gitlab.RegistryRepository, error) {
	// Cria uma lista de registros de repositórios no GitLab.
	registryRepos := []*gitlab.RegistryRepository{}

	// Cria as opções de lista de registros do GitLab para lidar com paginação e obter todos os resultados.
	opts := &gitlab.ListRegistryRepositoriesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 10,
		},
	}

	// Realiza a paginação para obter todos os registros de repositórios do projeto.
	for {
		result, resp, err := client.api.ContainerRegistry.ListGroupRegistryRepositories(groupPath, opts)
		if resp.StatusCode == http.StatusNotFound {
			// Se ocorrer um erro ao obter os registros, retorne um erro.
			return nil, fmt.Errorf("erro ao obter repositórios do registry do grupo %s: %v", groupPath, err)
		}
		if err != nil {
			return nil, err
		}

		// Adiciona o resultado à lista de registros de repositórios.
		registryRepos = append(registryRepos, result...)

		// Verifica se devemos interromper a paginação.
		if resp.NextPage == 0 {
			break
		}

		// Vá para a próxima página.
		opts.Page = resp.NextPage
	}

	// Retorna a lista de registros de repositórios.
	return registryRepos, nil
}
