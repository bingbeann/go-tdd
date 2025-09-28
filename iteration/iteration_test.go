package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatConcat(b *testing.B) {
	for b.Loop() {
		var s string
		for range 1_000 {
			s += "a"
		}
	}
}

func BenchmarkRepeatStringBuilder(b *testing.B) {
	for b.Loop() {
		var sb strings.Builder
		for range 1_000 {
			sb.WriteString("a")
		}
		_ = sb.String()
	}
}
