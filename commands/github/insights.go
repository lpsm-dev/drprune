package github

import (
	"bufio"
	"context"
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	"github.com/ci-monk/drprune/internal/constants"
	"github.com/ci-monk/drprune/internal/log"
	"github.com/ci-monk/drprune/internal/utils"
	gh "github.com/ci-monk/drprune/pkg/github"
	"github.com/google/go-github/v41/github"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var interactive bool

// NewCmdInsights handlers gh get insights command.
func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:   "insights",
		Short: "Get insights of GitHub Registry (ghcr.io)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(constants.ASCIInsights)

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
					PerPage: 100,
				},
			})
			if err != nil {
				log.Errorf("List package err: %v", err)
			}

			// Get total packages
			totalPackages, totalTaggedPackages, totalUntaggedPackages := len(pkgs), 0, 0

			// Check if user have packages.
			if len(pkgs) > 0 {
				// If user have packages, get each package content.
				for index, pkg := range pkgs {
					totalTagged, totalUntagged := 0, 0

					c := gh.ContainerPackage{
						ID:         int(*pkg.ID),
						Name:       *pkg.Name,
						Index:      index + 1,
						Owner:      *pkg.Owner.Login,
						Visibility: *pkg.Visibility,
						CreatedAt:  pkg.CreatedAt.Time,
					}

					c.PrettyPrintContainerPackage()

					// Get all versions of the package.
					pkgVersions, _, err := client.Users.PackageGetAllVersions(ctx, name, "container", utils.EncodeParam(pkg.GetName()), nil)
					if err != nil {
						log.Fatal(err)
					}

					pterm.DefaultSection.Println("Package Versions Information")

					// Loop each version of the package.
					for _, pkgVersion := range pkgVersions {
						fmt.Printf("> Package version name: %s\n", *pkgVersion.Name)
						fmt.Printf("> Package version id: %d\n", *pkgVersion.ID)
						tags := pkgVersion.GetMetadata().GetContainer().Tags
						if len(tags) == 0 {
							totalUntagged++
							fmt.Printf("> Empty tags\n\n")
							fmt.Println()
							continue
						}
						fmt.Printf("> Package tags: %v\n\n", tags)
						for i := 0; i < len(tags); i++ {
							totalTagged++
						}
					}

					pterm.DefaultSection.Println("Resume Information")

					totalTaggedPackages += totalTagged
					totalUntaggedPackages += totalUntagged
					fmt.Printf("----> Total images tagged for %s package: %v", pkg.GetName(), totalTagged)
					fmt.Printf("\n----> Total images untagged for %s package: %v\n\n", pkg.GetName(), totalUntagged)

					if interactive {
						consoleReader := bufio.NewReaderSize(os.Stdin, 1)
						fmt.Print("\n> Press [ENTER] to see the next or [ESC] to exit: ")
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
				// If user no gave packages
				log.Warnf("The user %s don't any container packages!", name)
			}

			pterm.Println()
			pterm.DefaultHeader.
				WithBackgroundStyle(pterm.NewStyle(pterm.BgDarkGray)).
				WithTextStyle(pterm.NewStyle(pterm.FgLightYellow)).
				WithMargin(45).
				Println("Final Information")
			pterm.Println()

			versionTable := table.NewWriter()
			versionTable.SetStyle(table.StyleLight)
			versionTable.Style().Options.DrawBorder = true
			versionTable.SetOutputMirror(os.Stdout)
			versionTable.SetAllowedRowLength(80)
			versionTable.AppendHeader(table.Row{"Info", "Content"})
			versionTable.AppendRows([]table.Row{
				{"➤ Total tagged images", totalTaggedPackages},
				{"➤ Total untagged images", totalUntaggedPackages},
				{"➤ Total packages", totalPackages},
			})
			versionTable.Render()
			pterm.Println()
		},
	}

	insightsCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Interactive user mode")
	return insightsCmd
}
