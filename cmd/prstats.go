package cmd

import (
	"devstats/configs"
	"devstats/internal/adapter"
	"devstats/internal/adapter/crucible"
	"devstats/internal/adapter/github"
	"devstats/internal/client"
	"devstats/internal/display"
	"devstats/internal/service/reviewstats"
	"devstats/internal/stats"
	"devstats/internal/utils"
	gogithub "github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"time"
)

func init() {
	rootCmd.AddCommand(prstatsCmd)
}

var prstatsCmd = &cobra.Command{
	Use:   "prstats",
	Short: "Get PR stats.",
	Long:  `Get PR stats.`,
	Run: func(cmd *cobra.Command, args []string) {
		display.StandardGreeting()

		appConfig, err := configs.Read()
		utils.CheckPrintError(err)

		ctx := context.Background()

		crucibleClient := client.NewClient()
		githubClient, err := githubClientFactory(ctx, appConfig.Github.Access.BaseUrl, appConfig.Github.Access.Key)
		utils.CheckFatalError(err)

		startDate := utils.NewStatsTime(time.Now().AddDate(0, 0, -28))
		endDate := utils.NewStatsTime(time.Now())

		adapters := append(
			githubAdapterFactory(githubClient, appConfig.Github.Repositories),
			crucibleAdapterFactory(crucibleClient, appConfig.Crucible)...,
		)

		statsChannel := reviewstats.RunPipeline(ctx, startDate, endDate, adapters, stats.GetAverageOpenTime)
		for statsGroup := range statsChannel {
			for key, value := range statsGroup {
				display.PrOpenStatsDisplay(key, value)
			}
		}
	},
}

func githubAdapterFactory(ghClient *gogithub.Client, repositories []configs.RepositoryOptions) []adapter.ReviewAdapter {
	var adapters []adapter.ReviewAdapter
	for _, repo := range repositories {
		adapters = append(adapters, github.Adapter(ghClient, repo, github.ReviewsToReviewList))
	}

	return adapters
}

func crucibleAdapterFactory(httpClient client.HttpClientWrapper, config configs.CrucibleConfiguration) []adapter.ReviewAdapter {
	var adapters []adapter.ReviewAdapter
	for _, project := range config.Projects {
		adapters = append(adapters, crucible.Adapter(httpClient, config.Access, project, crucible.ReviewsToReviewList))
	}

	return adapters
}

func githubClientFactory(ctx context.Context, url string, accessKey string) (*gogithub.Client, error) {
	httpClient := client.NewClientWithOauth(accessKey, ctx)
	return github.NewGithubClient(httpClient, url)
}
