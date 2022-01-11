package gitlab

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var token string
var ns string
var url string

func NewCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gl",
		Short: "Perform gitlab operations",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gl images")
		},
	}
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", os.Getenv("GL_TOKEN"), "GitLab API Token (*)")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://gitlab.com/api/v4", "GitLab API URL")
	rootCmd.PersistentFlags().StringVarP(&ns, "ns", "n", "lpmatos", "GitLab Namespace - Group or Repo (*)")

	rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	return rootCmd
}
