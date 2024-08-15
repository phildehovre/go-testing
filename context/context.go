package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	// Anything that implements the fetch method
	// satisfies the Store interface
	// Store takes in the request context,
	// therefore propagating the pipeline
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			// LOG ERROR
			return
		}
		fmt.Fprint(w, data)
	}
}
