package racer

import "net/http"

func Racer(a, b string) (winner string) {

	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)

// }
// Channel is useful in this case

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
