package github

import (
	"github.com/google/go-github/github"
)

type RepositoryPullRequests struct {
	Repository   string
	PullRequests []*github.PullRequest
}
