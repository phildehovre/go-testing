package racer2

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {
// 	t.Run("will return faster url", func(t *testing.T) {

// 		fastUrl := "www.facebook.com"
// 		slowUrl := "www.quii.dev"

// 		want := slowUrl
// 		got := Racer(fastUrl, slowUrl)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

func TestRacer(t *testing.T) {
	t.Run("test website speed", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %q got %q", got, want)
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
