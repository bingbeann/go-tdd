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

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.duration += d
}

func TestCountdown(t *testing.T) {
	requireNoError := func(tb testing.TB, err error) {
		tb.Helper()

		if err != nil {
			tb.Fatalf("expected no error but got %q", err)
		}
	}

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		stubSleeper := &StubSleeper{}

		err := Countdown(buffer, 3, stubSleeper)
		requireNoError(t, err)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}

		err := Countdown(spyCountdownOperations, 3, spyCountdownOperations)
		requireNoError(t, err)

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

	t.Run("sleep for total 3 seconds", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyTime := &SpyTime{}

		err := Countdown(buffer, 3, spyTime)
		requireNoError(t, err)

		got := spyTime.duration
		want := 3 * time.Second

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
