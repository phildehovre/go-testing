package racer2

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureDuration(a)
	bDuration := measureDuration(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}
