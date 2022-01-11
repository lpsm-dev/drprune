package github

import (
	"context"

	log "github.com/lpmatos/drprune/internal/log"
	"github.com/lpmatos/drprune/internal/utils"

	"github.com/google/go-github/v41/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var dryRun, untagged bool
var olderThan int

func NewCmdImages() *cobra.Command {
	var imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "Perform prune images operation on GitHub Registry (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			container = utils.EncodeParam(container)

			// Auth in github client
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			/*
				        Getting all versions a package
							    * name      - GitHub Username
							    * container - GitHub Container/Package Name
			*/
			pkgVersions, _, err := client.Users.PackageGetAllVersions(ctx, name, "container", container)
			if err != nil {
				log.Fatal(err)
			}

			size := len(pkgVersions)

			log.Info(pkgVersions)
			log.Infof("Package: %s", utils.DecodeParam(container))
			log.Infof("We have %v versions in this package", size)

			// Loop in the list of package versions
			for _, pkg := range pkgVersions {
				// Gettings packages tags - f
				// Implement filter by date
				tags := pkg.Metadata.Container.Tags

				// Implement deletion recursive tags

				// Check untagged
				if !(len(tags) == 0) {
					// Is tagged
					continue
				}

				version := pkg.GetID()

				// check dry-run

				resp, err := client.Users.PackageDeleteVersion(ctx, name, "container", container, version)
				if err != nil {
					log.Fatal(err)
				}
				log.Debug(resp)

				/*if resp.StatusCode == 400 {
				    log.Warnf("Oh no bro... somethings is wrong. We can't delete version [ %v ] of %v.", container)
				  } else if resp.StatusCode == 404 {
				    log.Warnf("%v package %v version [%v] not found", name, pkg, container)
				  } else {
				    log.Info("start prune %v package %v version %v, err: %v", name, pkg, container, err)
				  }*/

			}

		},
	}
	imagesCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Controlling whether to execute the action as a dry-run")
	imagesCmd.PersistentFlags().BoolVarP(&untagged, "untagged", "u", false, "Boolean controlling whether untagged versions should be pruned")
	imagesCmd.PersistentFlags().IntVarP(&olderThan, "older-than", "o", 0, "Minimum age in days of a version before it is pruned")
	return imagesCmd
}
