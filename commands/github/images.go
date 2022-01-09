package github

import (
	"context"
	"fmt"

	log "github.com/lpmatos/drprune/internal/log"

	"github.com/google/go-github/v41/github"
	"github.com/lpmatos/drprune/internal/utils"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func NewCmdImages() *cobra.Command {
	var imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "Perform prune old images on GitHub (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Check flags
			if utils.IsEmpty(token) {
				log.Fatal("Error set token")
			}
			if utils.IsEmpty(container) {
				log.Fatal("Error set container")
			}

			// Auth in github client
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			// Getting all versions a package
			pkgVersions, _, err := client.Users.PackageGetAllVersions(ctx, name, "container", container)
			if err != nil {
				log.Fatal(err)
			}

			// Check if we have versions in this package
			if len(pkgVersions) == 0 {
				log.Fatalf("We don't have versions in this package...")
			} else {
				log.Info(pkgVersions)
				log.Infof("We have %s", len(pkgVersions))
				// Loop in the list of package versions
				for _, pkg := range pkgVersions {
					tags := pkg.GetMetadata().Container.Tags
					// filter by date
					// check untagged
					if !(len(tags) == 0) {
						continue
					}
					created, version := pkg.CreatedAt, pkg.GetID()
					fmt.Println(created, version)
					// check dry-run
					response, err := client.Users.PackageDeleteVersion(ctx, name, "container", container, version)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Delete Success!", response)
					fmt.Println("-----------------")
				}
			}
		},
	}
	return imagesCmd
}
