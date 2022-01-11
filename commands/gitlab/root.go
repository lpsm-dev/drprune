package gitlab

import (
	"fmt"
	"os"

	log "github.com/lpmatos/drprune/internal/log"
	"github.com/spf13/cobra"
)

var token, ns, url string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gl",
		Short: "Perform gitlab operations",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gl images")
		},
	}
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitLab API Token (*)")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://gitlab.com/api/v4", "GitLab API URL")
	rootCmd.PersistentFlags().StringVarP(&ns, "ns", "n", "lpmatos", "GitLab Namespace - Group or Repo (*)")

	rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	return rootCmd
}

func checkCmdParams() {
	token = os.Getenv("GL_TOKEN")
	url = os.Getenv("GL_URL")
	ns = os.Getenv("GL_NAMESPACE")

	if token == "" {
		log.Fatalln("Please, set a GitLab Token")
	}

	if url == "" {
		log.Fatalln("Please, set a GitLab URL")
	}

	if ns == "" {
		log.Fatalln("Please, set a GitLab Namespace")
	}
}
