package server

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

func TestServer(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		want := "hello, world"
		store := &SpyStore{want, false}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		got := resp.Body.String()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if store.cancelled {
			t.Error("store should not have cancelled")
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		want := "hello, world"
		store := &SpyStore{want, false}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
