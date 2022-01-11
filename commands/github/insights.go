package github

import (
	"bufio"
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	"github.com/google/go-github/v41/github"
	"github.com/lpmatos/drprune/internal/log"
	"github.com/lpmatos/drprune/internal/utils"
	gh "github.com/lpmatos/drprune/pkg/github"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var interactive bool

func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitHub Registry (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			totalPackages := 0
			container = utils.EncodeParam(container)

			client, ctx, err := gh.NewClient(token)
			if err != nil {
				log.Fatal(err)
			}

			/*releases, _, err := client.Repositories.ListReleases(ctx, name, "loli", &github.ListOptions{PerPage: 100})
			if err != nil {
				log.Fatal(err)
			}

			lastRelease := releases[0]

			var (
				osMap = map[string]string{
					"darwin":  "Darwin",
					"linux":   "Linux",
					"windows": "Windows",
				}

				archMap = map[string]string{
					"386":   "i386",
					"amd64": "x86_64",
					"arm":   "arm",
				}
			)

			var (
				OS   = osMap[runtime.GOOS]
				ARCH = archMap[runtime.GOARCH]
			)

			log.Infoln(OS)
			log.Infoln(ARCH)
			for _, i := range lastRelease.Assets {
				name := i.GetName()
				if strings.Contains(name, OS) {
					if strings.Contains(name, ARCH) {
						log.Debugln(i.GetBrowserDownloadURL())
					}
				}
			}

			os.Exit(1)
			*/

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

			// Check if user have packages.
			if len(pkgs) > 0 {
				// If user have packages, loop.
				for _, pkg := range pkgs {
					fmt.Printf("\n\n=================================================\n\n")

					totalTags, totalUntagged := 0, 0

					c := gh.ContainerPackage{
						ID:         int(*pkg.ID),
						Name:       *pkg.Name,
						Owner:      *pkg.Owner.Login,
						Visibility: *pkg.Visibility,
						CreatedAt:  pkg.CreatedAt.Time,
					}

					c.PrettyPrintContainerPackage()

					// Get all versions of the package.
					pkgVersions, _, err := client.Users.PackageGetAllVersions(ctx,
						name, "container", utils.EncodeParam(pkg.GetName()),
					)
					if err != nil {
						log.Fatal(err)
					}

					pterm.Println()
					pterm.DefaultSection.Println("Package Versions Information")

					// Loop each version of the package.
					for _, pkgVersion := range pkgVersions {
						fmt.Printf("\n> Package version name: %s\n", *pkgVersion.Name)
						fmt.Printf("> Package version id: %d\n", *pkgVersion.ID)

						tags := pkgVersion.GetMetadata().GetContainer().Tags
						if len(tags) == 0 {
							totalUntagged++
							fmt.Printf("> Empty tags\n")
							continue
						}

						fmt.Printf("> Package tags: %v\n", tags)
						for i := 0; i < len(tags); i++ {
							totalTags++
						}
					}

					pterm.Println()
					pterm.DefaultSection.Println("Resume Information")
					fmt.Printf("----> Total tagged for %s package: %v", pkg.GetName(), totalTags)
					fmt.Printf("\n----> Total untagged for %s package: %v", pkg.GetName(), totalUntagged)
					totalPackages++

					if interactive {
						consoleReader := bufio.NewReaderSize(os.Stdin, 1)
						fmt.Print("\n\n> Press [ENTER] to see the next or [ESC] to exit: ")
						input, _ := consoleReader.ReadByte()
						ascii := input
						if ascii == 27 {
							fmt.Printf("\nExiting...\n")
							os.Exit(0)
						}

						tm.Clear()
						tm.MoveCursor(1, 1)
						tm.Flush()
					}
				}
			} else {
				log.Warnf("The user %s don't any tags!", name)
			}

			fmt.Printf("\n\n=================================================\n")
			pterm.Println()
			pterm.DefaultSection.Println("Final Information")
			fmt.Printf("\n----> The user %s have %v packages\n", name, totalPackages)
		},
	}

	insightsCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Interactive user mode")
	return insightsCmd
}
