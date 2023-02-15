package gitlab

import (
	"os"

	log "github.com/ci-monk/drprune/internal/log"
	"github.com/spf13/cobra"
)

var token, ns, url string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gl",
		Short: "Perform GitLab operations",
		Long:  ``,
	}

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitLab API Token - $GL_TOKEN (*)")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "GitLab API URL - $GL_API_URL")
	rootCmd.PersistentFlags().StringVarP(&ns, "namespace", "n", "", "GitLab Namespace - $GL_NAMESPACE (*)")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())

	return rootCmd
}

func checkCmdParams() {
	token = os.Getenv("GL_TOKEN")
	url = os.Getenv("GL_API_URL")
	ns = os.Getenv("GL_NAMESPACE")

	if token == "" {
		log.Fatalln("Please set a GitLab Token")
	}

	if url == "" {
		log.Fatalln("Please set a GitLab URL")
	}

	if ns == "" {
		log.Fatalln("Please set a GitLab Namespace")
	}
}
