package countdown

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	err := Countdown(buffer, 3, spySleeper)
	if err != nil {
		t.Fatalf("failed to countdown: %v", err)
	}

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, wanted 3 but got %d", spySleeper.Calls)
	}
}
