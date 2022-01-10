package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
	"github.com/lpmatos/drprune/internal/log"
	"github.com/lpmatos/drprune/internal/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

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
					totalTags, totalUntagged := 0, 0
					fmt.Printf("\n\n=================================================\n\n")
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
						fmt.Printf("\n> Package version: %s\n", *pkgVersion.Name)
						tags := pkgVersion.GetMetadata().GetContainer().Tags
						if len(tags) == 0 {
							totalUntagged++
						}
						fmt.Printf("> Package tags: %v\n", tags)
						for _, tag := range tags {
							log.Debugf("%v\n", tag)
							totalTags++
						}
					}

					pterm.Println()
					pterm.DefaultSection.Println("Resume Information")
					fmt.Printf("----> Total tagged for %s package: %v", pkg.GetName(), totalTags)
					fmt.Printf("\n----> Total untagged for %s package: %v", pkg.GetName(), totalUntagged)
					totalPackages++
				}
			} else {
				log.Warnf("The user %s don't any tags!", name)
			}

			fmt.Printf("\n\n=================================================\n")
			pterm.Println()
			pterm.DefaultSection.Println("Final Information")
			fmt.Printf("\n----> Total of packages for %s user: %v\n", name, totalPackages)
		},
	}
	return insightsCmd
}
