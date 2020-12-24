package github

import (
	"context"
	"devstats/configs"
	"devstats/internal/adapter"
	"devstats/internal/client"
	"devstats/internal/model"
	"devstats/internal/utils"
	"github.com/google/go-github/github"
	"time"
)

func NewGithubClient(client client.HttpClientWrapper, url string) (*github.Client, error) {
	return github.NewEnterpriseClient(
		url,
		"",
		client.Client,
	)
}

func Adapter(
	githubClient *github.Client,
	repository configs.RepositoryOptions,
	transformer ReviewTransformer,
) adapter.ReviewAdapter {
	return func(
		ctx context.Context,
		startDate utils.StatsTime,
		endDate utils.StatsTime,
	) model.RepositoryReviews {

		options := github.PullRequestListOptions{
			State:     "closed",
			Base:      repository.Branch,
			Sort:      "closed",
			Direction: "desc",
			ListOptions: github.ListOptions{
				Page:    0,
				PerPage: 100,
			},
		}

		reviews := fetch(repository, ctx, startDate.Time, githubClient, options, []*github.PullRequest{})

		transformedReviews, err := transformer(repository.Repository, reviews)
		utils.CheckFatalError(err) // TODO: let's be a little more fault tolerant

		return transformedReviews
	}
}

func fetch(
	repoOptions configs.RepositoryOptions,
	ctx context.Context,
	stopDate time.Time,
	client *github.Client,
	options github.PullRequestListOptions,
	prs []*github.PullRequest,
) []*github.PullRequest {

	pullRequests, _, err := client.PullRequests.List(
		ctx,
		repoOptions.Owner,
		repoOptions.Repository,
		&options)
	utils.CheckPrintError(err)

	for _, pr := range pullRequests {
		if pr.ClosedAt.After(stopDate) {
			prs = append(prs, pr)
		}
	}

	if options.ListOptions.PerPage*(options.ListOptions.Page+1) == len(prs) {
		options.ListOptions.Page++
		return fetch(repoOptions, ctx, stopDate, client, options, prs)
	}

	return prs
}
