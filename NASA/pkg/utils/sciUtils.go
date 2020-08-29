package utils

import (
	"time"
)

//TimeDelta calculate time delta.
func TimeDelta(Hour int64, Minute int64, Second float64) int {
	baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	date := baseTime.Add(time.Duration(Hour)*time.Hour + time.Duration(Minute)*time.Hour + time.Duration(Second)*time.Second)
	return (date.Second())
}
