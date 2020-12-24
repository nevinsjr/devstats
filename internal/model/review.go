package model

import (
	"time"
)

type RepositoryReviews struct {
	Repository string
	Reviews    []Review
}

type Review struct {
	Id          string
	CreatedDate time.Time
	ClosedDate  time.Time
	ReviewLink  string
}
