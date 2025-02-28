package metrics

import (
	"time"
)

func Seconds(duration time.Duration) float64 {
	return float64(duration) / float64(time.Second)
}

func SinceSeconds(started time.Time) float64 {
	return float64(time.Since(started)) / float64(time.Second)
}

func IsError(err error) string {
	if err == nil {
		return "0"
	}
	return "1"
}
