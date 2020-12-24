package display

import (
	"devstats/internal/model"
	"fmt"
)

// TODO: this should be a little more abstract or live in another package
func PrOpenStatsDisplay(repository string, stats model.ReviewStats) {
	fmt.Printf(
		"Average PR Open Time for the %s Repository\n%s %s %s\n\n",
		repository, stats.AvgOpenTime.Days, stats.AvgOpenTime.Hours, stats.AvgOpenTime.Minutes)
	fmt.Printf("Outliers: %d\n\n", len(stats.Outliers))
	for pr, time := range stats.Outliers {
		fmt.Printf("PR ID: %s\n%s, %s, %s\n\n", pr, time.Days, time.Hours, time.Minutes)
	}
}

func PrCommentsDisplay(stats model.ReviewComments) {
	println("\nAverage PR Comments:  " + stats.Average)
}
