package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
	"github.com/lpmatos/drprune/internal/log"
	"github.com/lpmatos/drprune/internal/utils"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

/*
Considerations:

- A user can have a list of packages.
  - Each package have a type: [container, maven, npm].
  - Each package have versions.
    - Each version of a package can have a name.
    - Each version of a package can be tagged or not.
*/

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitHub Registry (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			container = utils.EncodeParam(container)
			totalPackages := 0

			// Auth in github client
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			// Get all packages of user.
			pkgs, _, err := client.Users.ListPackages(ctx, name, &github.PackageListOptions{
				PackageType: github.String("container"),
				ListOptions: github.ListOptions{
					PerPage: 300,
				},
			})
			if err != nil {
				log.Errorf("List package err: %v", err)
			}

			// Check if user have package
			if len(pkgs) > 0 {
				// If user have a package, loop each package.
				for _, pkg := range pkgs {
					fmt.Println("=================================================")
					fmt.Printf("> Package name: %v\n", pkg.GetName())

					// Get all versions of the package.
					pkgVersions, _, err := client.Users.PackageGetAllVersions(ctx,
						name, "container", utils.EncodeParam(pkg.GetName()),
					)
					if err != nil {
						log.Fatal(err)
					}

					// Loop each version of the package.
					for _, pkgVersion := range pkgVersions {
						fmt.Printf("> Package version: %s\n", *pkgVersion.Name)
						tags := pkgVersion.GetMetadata().GetContainer().Tags
						fmt.Printf("> Package tags: %v\n", tags)
						for _, tag := range tags {
							log.Debugf("%v\n", tag)
						}
					}
					totalPackages++
				}
			} else {
				log.Warnf("The user %s don't any tags!", name)
			}

			log.Debugf("\n\nTotal packages: %v", totalPackages)
		},
	}
	return insightsCmd
}
