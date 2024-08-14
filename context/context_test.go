package context

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func NewSpyStore() *SpyStore {
	return &SpyStore{response: "hello, world"}
}

func TestServer(t *testing.T) {
	t.Run("server is running", func(t *testing.T) {
		// create dummy store
		spy := NewSpyStore()
		// create server with dummy store
		svr := Server(spy)

		// send test GET request to "/" endpoint
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// The recorder implements the ResponseWriter interface
		// and will capture the server's response
		response := httptest.NewRecorder()

		// initialize the server,it returns a basic HandlerFunc
		svr.ServeHTTP(response, request)

		// This assertion will compare the server response to our spy response
		if response.Body.String() != spy.response {
			t.Errorf("got %s, want %s", response.Body.String(), spy.response)
		}

	})

	// This test will call upon the WithCancel method to graft
	// cancelling functionality onto our request context
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		spy := NewSpyStore()
		spy.response = "hello world"
		spy.t = t
		svr := Server(spy)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Creates a new instance of context with cancel tooling
		// gives us access to the cancel func
		cancellingCtx, cancel := context.WithCancel(request.Context())

		// will calls cancel after countdown
		// This simulates user cancellation
		time.AfterFunc(5*time.Millisecond, cancel)

		// Grafts cancel context onto our request context and
		request = request.WithContext(cancellingCtx)
		fmt.Println("REQUEST WITH CANCELLING CONTEXT: ", request.Context())
		response := &SpyResponseWriter{}

		// Initialize server with response writer and request with context
		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}
	})

	t.Run("returns data from the store", func(t *testing.T) {

		data := "hello world"
		store := NewSpyStore()
		store.response = data

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
}
