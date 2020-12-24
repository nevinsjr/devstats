package utils

import (
	"strconv"
	"strings"
	"time"
)

var dayDuration = defineDayDuration()
var weekDuration = dayDuration * 7

type StatsTime struct {
	Time     time.Time
	UnixTime int64
}

func NewStatsTime(time time.Time) StatsTime {
	return StatsTime{
		time,
		time.UnixNano() / 1000000,
	}
}

func defineDayDuration() time.Duration {
	dayDuration, _ := time.ParseDuration("24h")
	return dayDuration
}

func IsActivityGreaterThanOneWeek(duration time.Duration) bool {
	return duration > weekDuration
}

func IsActivityOverSingleWeekend(from time.Weekday, to time.Weekday) bool {
	return IsActivityOverWeekend(from, to) &&
		!IsActivityOnWeekend(from) &&
		!IsActivityOnWeekend(to)
}

func IsActivityOverWeekend(from time.Weekday, to time.Weekday) bool {
	return to < from
}

func IsActivityOnWeekend(day time.Weekday) bool {
	return day == 6 || day == 0
}

func ReduceDurationByWeekend(duration time.Duration) time.Duration {
	return duration - dayDuration*2
}

func DeriveWeekendsFromDuration(duration time.Duration) int {
	// TODO: this is a little ridiculous
	weekends, _ := strconv.Atoi(
		strings.TrimSuffix((duration / time.Hour / (24 * 7)).Round(weekDuration/time.Hour/(24*7)).String(), "ns"))
	return weekends
}

func GetAverageDurationByPeriod(duration time.Duration, period int) time.Duration {
	return duration / time.Duration(period)
}
