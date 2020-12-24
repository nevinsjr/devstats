package adapter

import (
	"context"
	"devstats/internal/model"
	"devstats/internal/utils"
)

type ReviewAdapter func(
	ctx context.Context,
	startDate utils.StatsTime,
	endDate utils.StatsTime,
) model.RepositoryReviews
