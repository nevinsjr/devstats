package stats

import (
	"devstats/internal/model"
	"devstats/internal/utils"
	"strings"
	"time"
)

func GetAverageOpenTime(reviews []model.Review) model.ReviewStats {

	totalDuration := time.Duration(0)
	outliers := make(map[string]model.ReviewOpenTime)
	for _, review := range reviews {

		reviewDuration := review.ClosedDate.Sub(review.CreatedDate)
		from := review.CreatedDate.Weekday()
		to := review.ClosedDate.Weekday()

		reviewDuration, isModified := checkHandleActivityGreaterThanOneWeek(reviewDuration)
		if isModified {
			totalDuration += reviewDuration

			// TODO: shameful copy pasta
			days := reviewDuration / time.Hour / 24
			hours := (reviewDuration - (days * 24 * time.Hour)) / time.Hour
			minutes := (reviewDuration - ((days * 24 * time.Hour) + (hours * time.Hour))) / time.Minute

			// store PRs greater than one week as an outlier
			outliers[review.Id] = model.ReviewOpenTime{
				Days:    strings.TrimSuffix(days.String(), "ns") + " Days",
				Hours:   strings.TrimSuffix(hours.String(), "ns") + " Hours",
				Minutes: strings.TrimSuffix(minutes.String(), "ns") + " Minutes",
			}

			continue
		}

		reviewDuration, isModified = checkHandleActivityOverSingleWeekend(reviewDuration, from, to)
		if isModified {
			totalDuration += reviewDuration
			continue
		}

		totalDuration += reviewDuration
	}

	averageDuration := utils.GetAverageDurationByPeriod(totalDuration, len(reviews))

	averageDays := averageDuration / time.Hour / 24
	averageHours := (averageDuration - (averageDays * 24 * time.Hour)) / time.Hour
	averageMinutes := (averageDuration - ((averageDays * 24 * time.Hour) + (averageHours * time.Hour))) / time.Minute

	return model.ReviewStats{
		AvgOpenTime: model.ReviewOpenTime{
			Days:    strings.TrimSuffix(averageDays.String(), "ns") + " Days",
			Hours:   strings.TrimSuffix(averageHours.String(), "ns") + " Hours",
			Minutes: strings.TrimSuffix(averageMinutes.String(), "ns") + " Minutes",
		},
		Outliers: outliers,
	}
}

// Handle activity greater than 1 week, which could span 1...n weekends
func checkHandleActivityGreaterThanOneWeek(prDuration time.Duration) (time.Duration, bool) {

	if utils.IsActivityGreaterThanOneWeek(prDuration) {
		weekends := utils.DeriveWeekendsFromDuration(prDuration)
		for i := 0; i < weekends; i++ {
			prDuration = utils.ReduceDurationByWeekend(prDuration)
		}
		return prDuration, true
	}

	return prDuration, false
}

// Handle activity over the weekend.  Work that was started or closed
// on a weekend day does not qualify for reduction.
func checkHandleActivityOverSingleWeekend(
	prDuration time.Duration,
	from time.Weekday,
	to time.Weekday,
) (time.Duration, bool) {

	if utils.IsActivityOverSingleWeekend(from, to) {
		prDuration = utils.ReduceDurationByWeekend(prDuration)
		return prDuration, true
	}

	return prDuration, false
}
