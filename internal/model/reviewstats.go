package model

type ReviewStats struct {
	AvgOpenTime ReviewOpenTime
	Outliers    map[string]ReviewOpenTime
}

type ReviewOpenTime struct {
	Days    string
	Hours   string
	Minutes string
}

type ReviewComments struct {
	Total   string
	Average string
}
