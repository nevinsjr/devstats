package crucible

import (
	"context"
	"devstats/configs"
	"devstats/internal/adapter"
	"devstats/internal/client"
	"devstats/internal/model"
	"devstats/internal/utils"
	"strconv"
)

func Adapter(
	httpClient client.HttpClientWrapper,
	access configs.Access,
	project configs.CrucibleProjectOptions,
	transformer ReviewTransformer,
) adapter.ReviewAdapter {
	return func(
		ctx context.Context,
		startDate utils.StatsTime,
		endDate utils.StatsTime,
	) model.RepositoryReviews {
		req, err := client.NewRequest("GET", access.BaseUrl, nil)
		req, err = req.WithOptions(
			client.BasicAuth(access.User, access.Key),
			client.Header("content-type", "application/json"),
			client.Header("accept", "application/json"),
			client.QueryParam("project", project.Name),
			client.QueryParam("fromDate", strconv.FormatInt(startDate.UnixTime, 10)),
			client.QueryParam("toDate", strconv.FormatInt(endDate.UnixTime, 10)),
			client.QueryParam("states", "Closed"),
		)
		utils.CheckFatalError(err)

		var crucibleReviews CrucibleProjectReviews
		err = httpClient.MakeRequest(req, client.Deserializer(&crucibleReviews))
		utils.CheckFatalError(err)

		transformedReviews, err := transformer(project.Name, crucibleReviews)
		utils.CheckFatalError(err) // TODO: let's be a little more fault tolerant
		return transformedReviews
	}
}
