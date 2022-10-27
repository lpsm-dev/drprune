package commands

import (
	githubCmd "github.com/ci-monk/drprune/commands/github"
	gitlabCmd "github.com/ci-monk/drprune/commands/gitlab"
	"github.com/ci-monk/drprune/internal/consts"
	"github.com/ci-monk/drprune/internal/debug"
	"github.com/ci-monk/drprune/internal/log"
	"github.com/ci-monk/drprune/internal/utils"
	"github.com/spf13/cobra"
)

var config = log.Config{}

var rootCmd = &cobra.Command{
	Use:   consts.BinaryName,
	Short: "Prune images store in GitHub/GitLab container registry",
	Long: `Description:

🤖 This is a pretty CLI you can use to prune docker images.
`,
	DisableSuggestions: false,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.SetDefault(config.Level, config.Format, config.Output, config.File, config.Verbose)
		err := log.Setup(
			log.WithConfig(config),
		)

		if err != nil {
			log.Warn("Error setting log: %v", err)
		}

		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			debug.Verbose = verbose
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&config.Level, "log-level", "debug", "set the logging level. One of: debug|info|warn|error")
	rootCmd.PersistentFlags().StringVar(&config.Format, "log-format", "color", "the formating of the logs. Available values: text|color|json|json-pretty")
	rootCmd.PersistentFlags().StringVar(&config.Output, "log-output", "stdout", "default log output. Available values: stdout|stderr|file")
	rootCmd.PersistentFlags().StringVar(&config.File, "log-file", utils.CreateLogFile("/var/log/drprune", "file"), "default CLI log file")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", true, "verbose output")
	rootCmd.AddCommand(githubCmd.NewCmd())
	rootCmd.AddCommand(gitlabCmd.NewCmd())
}
