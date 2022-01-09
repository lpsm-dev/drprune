package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

/*
- token: package:read and packages:delete
- organization: name of the organization, if empty the auth user are considered
- container: container package name
- dry-run: no delete
- older-than: vou limpar as images que exist a more de 7 dias.
- untagged
*/

func inside() {
	// create cobra cli

	ghToken := ""
	ghUser, ghPackageType, ghPackageName := "lpmatos", "container", "docker-crypto-miner"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	pkgVers, _, err := client.Users.PackageGetAllVersions(ctx, ghUser, ghPackageType, ghPackageName)
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgVers {
		tags := pkg.GetMetadata().Container.Tags
		// filter by date
		// check untagged
		if !(len(tags) == 0) {
			continue
		}
		created, version := pkg.CreatedAt, pkg.GetID()
		fmt.Println(created, version)
		// check dry-run
		response, err := client.Users.PackageDeleteVersion(ctx, ghUser, ghPackageType, ghPackageName, version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Delete Success!", response)
		fmt.Println("-----------------")
	}
}
