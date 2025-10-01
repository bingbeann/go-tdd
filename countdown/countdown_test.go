package countdown

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type StubSleeper struct{}

func (s *StubSleeper) Sleep(d time.Duration) {}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep(d time.Duration) {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return len(p), nil
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		stubSleeper := &StubSleeper{}

		err := Countdown(buffer, 3, stubSleeper)
		if err != nil {
			t.Fatalf("failed to countdown: %v", err)
		}

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}

		err := Countdown(spyCountdownOperations, 3, spyCountdownOperations)
		if err != nil {
			t.Fatalf("failed to countdown: %v", err)
		}

		got := spyCountdownOperations.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
