package reviewstats

import (
	"context"
	"devstats/internal/adapter"
	"devstats/internal/model"
	"devstats/internal/utils"
	"sync"
)

func RunPipeline(
	ctx context.Context,
	startDate utils.StatsTime,
	endDate utils.StatsTime,
	adapters []adapter.ReviewAdapter,
	statsBuilders ...func(reviews []model.Review) model.ReviewStats,
) <-chan map[string]model.ReviewStats {

	reviewChannels := getReviews(ctx, startDate, endDate, adapters)
	return compileStats(reviewChannels, statsBuilders...)
}

func getReviews(
	ctx context.Context,
	startDate utils.StatsTime,
	endDate utils.StatsTime,
	adapters []adapter.ReviewAdapter,
) []<-chan model.RepositoryReviews {

	var channels []<-chan model.RepositoryReviews

	for _, reviewAdapter := range adapters {
		channels = append(channels, executeAdapter(ctx, startDate, endDate, reviewAdapter))
	}

	return channels
}

func executeAdapter(
	ctx context.Context,
	startDate utils.StatsTime,
	endDate utils.StatsTime,
	reviewAdapter adapter.ReviewAdapter,
) <-chan model.RepositoryReviews {

	out := make(chan model.RepositoryReviews)

	go func(ra adapter.ReviewAdapter) {
		out <- ra(ctx, startDate, endDate)
		close(out)
	}(reviewAdapter)

	return out
}

func compileStats(
	in []<-chan model.RepositoryReviews,
	statsBuilders ...func(reviews []model.Review) model.ReviewStats,
) <-chan map[string]model.ReviewStats {

	var waitGroup sync.WaitGroup
	out := make(chan map[string]model.ReviewStats)

	var cumulativePrs []model.Review

	outputHandler := func(c <-chan model.RepositoryReviews) {
		for reviewGroup := range c {
			cumulativePrs = append(cumulativePrs, reviewGroup.Reviews...)
			for _, aggregator := range statsBuilders {
				out <- map[string]model.ReviewStats{reviewGroup.Repository: aggregator(reviewGroup.Reviews)}
			}
		}
		waitGroup.Done()
	}

	waitGroup.Add(len(in))
	for _, channel := range in {
		go outputHandler(channel)
	}

	// Wrap-up - emit cumulative stats and close
	go func() {
		waitGroup.Wait()
		for _, aggregator := range statsBuilders {
			out <- map[string]model.ReviewStats{"All": aggregator(cumulativePrs)}
		}
		close(out)
	}()

	return out
}
