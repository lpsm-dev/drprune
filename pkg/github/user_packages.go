package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v41/github"
)

type GithubClient struct {
	user string
	api  *github.Client
	ctx  context.Context
}

// GetAllContainerPackageVersions
func (gh *GithubClient) GetAllContainerPackageVersions(container string) ([]*github.PackageVersion, error) {
	var pkgVersions []*github.PackageVersion
	opts := &github.PackageListOptions{
		PackageType: github.String("container"),
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	for {
		results, resp, err := gh.api.Users.PackageGetAllVersions(gh.ctx, gh.user, "container", container, opts)
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("can't retrieve all versions of package, container %s/%s not found!: %v", gh.user, container, err)
		}
		if err != nil {
			return nil, fmt.Errorf("can't retrieve all versions of package: %v", err)
		}

		pkgVersions = append(pkgVersions, results...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return pkgVersions, nil
}
