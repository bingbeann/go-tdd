package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	mockServer := func(duration time.Duration) *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(duration)
			w.WriteHeader(http.StatusOK)
		}))
	}

	fastServer := mockServer(1 * time.Millisecond)
	slowServer := mockServer(2 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	got := Racer(slowServer.URL, fastServer.URL)
	want := fastServer.URL

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
