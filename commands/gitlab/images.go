package gitlab

import "github.com/spf13/cobra"

func NewCmdImages() *cobra.Command {
	var imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "Perform prune old images on GitLab (registry.gitlab.com)",
		Long:  ``,
	}
	return imagesCmd
}
