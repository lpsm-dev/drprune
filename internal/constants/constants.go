package constants

const (
	// BinaryName is the name of the app.
	BinaryName = "drprune"

	// DefaultTimestampFormat default time format.
	DefaultTimestampFormat = "2006-01-02_15:04:05"

	// ProjectURL is the project url of the app.
	ProjectURL = "https://github.com/lpmatos/drprune"

	// ReleaseURL default release URL.
	ReleaseURL = "https://api.github.com/repos/lpmatos/drprune/releases"
)

// CompletionHelpMessage - return the long description of completion command.
const CompletionHelpMessage = `To load completion for:

Bash:
- For bash, ensure you have bash completions installed and enabled.
- To access completions in your current shell, run.
- Alternatively, write it to a file and source in .bash_profile.
$ source <(drprune completion bash)

Zsh:
- For zsh, output to a file in a directory referenced by the $fpath shell.
$ source <(drprune completion zsh)
# To load completions for each session, execute once:
$ drprune completion zsh > "${fpath[1]}/_drprune"

Fish:
$ drprune completion fish | source
# To load completions for each session, execute once:
$ drprune completion fish > ~/.config/fish/completions/drprune.fish
`
