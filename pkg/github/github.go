package github

import (
	"context"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	username string         // GitHub Username or Organization name
	client   *github.Client // GitHub API
}

// NewClient returns a GithubClient struct and a error.
func NewClient(ctx context.Context, token, username, url string) (*GithubClient, error) {
	// Auth in github client - Return a Token sources.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	// Creates an *http.Client - Return a Token client.
	tc := oauth2.NewClient(ctx, ts)
	// Getting the GitHub client.
	client := github.NewClient(tc)

	// Populate GitHubClient struct.
	return &GithubClient{
		username: username,
		client:   client,
	}, nil
}
