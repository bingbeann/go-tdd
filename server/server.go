package server

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := store.Fetch(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if _, err = fmt.Fprint(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
