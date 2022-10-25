package gitlab

import "github.com/xanzy/go-gitlab"

type GitLabClient struct {
	api *gitlab.Client // GitLab API
}
