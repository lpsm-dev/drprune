package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ci-monk/drprune/internal/log"
	"github.com/google/go-github/v50/github"
)

// Get all packages of user
func (gh *GithubClient) GetAllContainerPackages(ctx context.Context, username string) ([]*github.Package, error) {
	pkgs, _, err := gh.client.Users.ListPackages(ctx, username, &github.PackageListOptions{
		PackageType: github.String("container"),
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao obter os pacotes de container do usuário %s: %v", username, err)
	}
	return pkgs, nil
}

// Get all container package versions off a user
func (gh *GithubClient) GetAllContainerPackagesVersions(ctx context.Context, container string) ([]*github.PackageVersion, error) {
	// Create a list of github package versions.
	pkgVersions := []*github.PackageVersion{}

	// Create a github package list options to handler pagination and get all results.
	opts := &github.PackageListOptions{
		PackageType: github.String("container"),
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	// Loop pagination
	for {
		// Return the GitHub all package versions of container.
		result, resp, err := gh.client.Users.PackageGetAllVersions(ctx, gh.username, "container", container, opts)
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("erro ao recuperar todas as versões do pacote %s/%s: %v", gh.username, container, err)
		}

		if err != nil {
			return nil, err
		}

		// Add result to the pkgVersion list.
		pkgVersions = append(pkgVersions, result...)

		// Check if we stop the pagination
		if resp.NextPage == 0 {
			break
		}

		// Go to the next page.
		opts.Page = resp.NextPage
	}

	// Return the list of packages.
	return pkgVersions, nil
}

// Delete container package version
func (gh *GithubClient) DeleteContainerPackageVersion(ctx context.Context, container string, pkg *github.PackageVersion, dryRun bool) {
	if dryRun {
		log.Infof("%d", pkg.GetID())
		log.Infof("%s", pkg.GetName())
		log.Infof("%s", pkg.GetMetadata().GetContainer().Tags)
	} else {
		resp, err := gh.client.Users.PackageDeleteVersion(ctx, gh.username, "container", container, pkg.GetID())
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("Sucessfull delete package %s on %s - Status %d", pkg.GetName(), container, resp.StatusCode)
	}
}
