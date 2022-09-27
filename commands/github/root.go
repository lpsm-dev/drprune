package github

import (
	"os"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/ci-monk/drprune/internal/utils"
	"github.com/spf13/cobra"
)

var token, name, container string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gh",
		Short: "Perform github operations",
		Long:  ``,
	}

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", os.Getenv("GH_TOKEN"), "GitHub API Token (*)")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", os.Getenv("GH_USERNAME"), "GitHub User/Organization Name (*)")
	rootCmd.PersistentFlags().StringVarP(&container, "container", "c", os.Getenv("GH_CONTAINER"), "GitHub Container Name (*)")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	return rootCmd
}

func checkCmdParams() {
	container = utils.EncodeParam(container)
	if token == "" {
		log.Fatalln("Please, set a GitHub Token")
	}

	if name == "" {
		log.Fatalln("Please, set a GitHub Name")
	}

	if container == "" {
		log.Fatalln("Please, set a GitHub Container Name")
	}
}
