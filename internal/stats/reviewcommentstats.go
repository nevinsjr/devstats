package stats

import (
	"devstats/internal/model"
	"github.com/google/go-github/github"
	"strconv"
)

func GetAverageComments(prs []*github.PullRequest) model.ReviewComments {

	totalComments := 0
	for _, pr := range prs {
		if pr.Comments != nil {
			totalComments += *pr.Comments
		}
	}

	return model.ReviewComments{
		Total:   strconv.Itoa(totalComments),
		Average: strconv.Itoa(totalComments / len(prs)),
	}
}
