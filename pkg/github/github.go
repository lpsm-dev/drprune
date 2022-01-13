package github

import (
	"context"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

// NewClient returns a GitHub client.
func NewClient(token, user string) (*GithubClient, error) {
	// Auth in github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GithubClient{
		user: user,
		api:  client,
		ctx:  ctx,
	}, nil
}
