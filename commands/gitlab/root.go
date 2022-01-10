package gitlab

import (
	"fmt"

	"github.com/spf13/cobra"
)

var token string
var ns string

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
	rootCmd.PersistentFlags().StringVarP(&ns, "ns", "n", "lpmatos", "GitLab Namespace - Group or Repo (*)")

	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("ns")

	rootCmd.AddCommand(NewCmdImages())
	rootCmd.AddCommand(NewCmdInsights())
	return rootCmd
}
