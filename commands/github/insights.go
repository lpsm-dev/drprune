package github

import (
	"bufio"
	"context"
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	"github.com/ci-monk/drprune/internal/consts"
	"github.com/ci-monk/drprune/internal/log"
	"github.com/ci-monk/drprune/internal/utils"
	gh "github.com/ci-monk/drprune/pkg/github"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var interactive bool

// NewCmdInsights handlers gh get insights command.
func NewCmdInsights() *cobra.Command {
	var insightsCmd = &cobra.Command{
		Use:     "insights",
		Aliases: []string{"i"},
		Short:   "Get insights of GitHub Registry (ghcr.io)",
		Long:    ``,
		Run: func(cmd *cobra.Command, args []string) {
			runInsights()
		},
	}
	insightsCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Interactive user mode")
	return insightsCmd
}

func runInsights() {
	fmt.Printf(consts.ASCIInsights)

	// Auth in github client
	ctx := context.Background()
	client, err := gh.NewClient(ctx, token, username, "")
	if err != nil {
		log.Fatal(err)
	}

	// Get all packages of user
	pkgs, err := client.GetAllContainerPackages(ctx, username)
	if err != nil {
		log.Errorf("List all user packages: %v", err)
	}

	// Get total packages
	totalPackages, totalTaggedPackages, totalUntaggedPackages := len(pkgs), 0, 0

	// Check if user have packages
	if len(pkgs) > 0 {
		// If user have packages, get each package content
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
			c.PrintTable()

			// Get all versions of the package
			pkgVersions, err := client.GetAllContainerPackagesVersions(ctx, utils.EncodeParam(pkg.GetName()))
			if err != nil {
				log.Errorf("List all container package versions: %v", err)
			}

			pterm.DefaultSection.Println("Package Versions Information")

			// Loop each version of the package
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
			fmt.Printf("> Total images tagged for %s package: %v", pkg.GetName(), totalTagged)
			fmt.Printf("\n> Total images untagged for %s package: %v\n\n", pkg.GetName(), totalUntagged)

			// If enable interactive mode
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
		// If user don't have container packages
		log.Warnf("The user %s don't any container packages!", username)
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
}
