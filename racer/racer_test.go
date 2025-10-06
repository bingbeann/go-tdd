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

	t.Run("get fastest url", func(t *testing.T) {
		fastServer := mockServer(1 * time.Millisecond)
		slowServer := mockServer(2 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		got, err := Racer(slowServer.URL, fastServer.URL)
		want := fastServer.URL

		if err != nil {
			t.Fatalf("expected no error but received one: %v", err)
		}

		if got != want {
			t.Errorf("want %q but got %q", want, got)
		}
	})

	t.Run("timeout after 10 seconds", func(t *testing.T) {
		server := mockServer(5 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 1*time.Millisecond)

		if err == nil {
			t.Fatal("expected one error but received none")
		}
	})
}
