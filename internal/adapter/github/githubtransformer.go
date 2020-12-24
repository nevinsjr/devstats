package github

import (
	"devstats/internal/model"
	"github.com/google/go-github/github"
	"strconv"
	"strings"
)

type ReviewTransformer func(repository string, pullRequests []*github.PullRequest) (model.RepositoryReviews, error)

func ReviewsToReviewList(repository string, pullRequests []*github.PullRequest) (model.RepositoryReviews, error) {
	var reviews []model.Review
	for _, pr := range pullRequests {
		review := model.Review{
			Id:          strconv.Itoa(*pr.Number),
			CreatedDate: *pr.CreatedAt,
			ClosedDate:  *pr.ClosedAt,
			ReviewLink:  *pr.URL,
		}
		reviews = append(reviews, review)
	}

	repositoryReviews := model.RepositoryReviews{
		Repository: strings.ToUpper(repository),
		Reviews:    reviews,
	}

	return repositoryReviews, nil
}
