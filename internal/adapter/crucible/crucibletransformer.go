package crucible

import (
	"devstats/internal/model"
	"strings"
	"time"
)

type ReviewTransformer func(project string, rawModel CrucibleProjectReviews) (model.RepositoryReviews, error)

func ReviewsToReviewList(project string, rawModel CrucibleProjectReviews) (model.RepositoryReviews, error) {
	var reviews []model.Review
	for _, crucibleReview := range rawModel.DetailedReviewData {
		createdDate, err := time.Parse(CrucibleTimeFormat, crucibleReview.CreateDate)
		if err != nil {
			return model.RepositoryReviews{Repository: ""}, err
		}

		closedDate, err := time.Parse(CrucibleTimeFormat, crucibleReview.CloseDate)
		if err != nil {
			return model.RepositoryReviews{Repository: ""}, err
		}

		review := model.Review{
			Id:          crucibleReview.PermaId.Id,
			CreatedDate: createdDate,
			ClosedDate:  closedDate,
			ReviewLink:  ReviewLinkBase + crucibleReview.PermaId.Id,
		}
		reviews = append(reviews, review)
	}

	repositoryReviews := model.RepositoryReviews{
		Repository: strings.ToUpper(project),
		Reviews:    reviews,
	}

	return repositoryReviews, nil
}
