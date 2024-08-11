package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func NewSpyStore() *SpyStore {
	return &SpyStore{response: "hello, world", cancelled: false}
}

func TestServer(t *testing.T) {
	t.Run("server is running", func(t *testing.T) {
		spy := NewSpyStore()
		svr := Server(spy)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != spy.response {
			t.Errorf("got %s, wat %s", response.Body.String(), spy.response)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		spy := NewSpyStore()
		svr := Server(spy)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != spy.response {
			t.Error("store was not told to cancel")
		}

	})
}
