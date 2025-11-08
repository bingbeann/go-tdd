package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		want := "hello, world"
		svr := Server(&SpyStore{want, false})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		got := resp.Body.String()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
