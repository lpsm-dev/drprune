package github

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

// GithubClient é um cliente para o GitHub.
type GithubClient struct {
	client   *github.Client // API do GitHub
	username string         // Nome de usuário ou nome da organização do GitHub
}

// NewClient retorna uma nova instância de GithubClient e um erro, se houver.
func NewClient(ctx context.Context, token, username, url string) (*GithubClient, error) {
	// Autenticação no cliente do GitHub - retorna uma fonte de token.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	// Cria um *http.Client - retorna um cliente de token.
	tc := oauth2.NewClient(ctx, ts)

	// Preenche a estrutura GithubClient.
	return &GithubClient{
		username: username,
		client:   github.NewClient(tc),
	}, nil
}
