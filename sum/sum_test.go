package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{1, 2, 3})
	want := []int{2, 5}

	if !slices.Equal(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func BenchmarkSliceAllocZero(b *testing.B) {
	for b.Loop() {
		s := make([]int, 0)
		for range 10 {
			s = append(s, 1)
		}

		_ = s
	}
}

func BenchmarkSliceAllocRequired(b *testing.B) {
	for b.Loop() {
		s := make([]int, 10)
		for i := range 10 {
			s[i] = 1
		}

		_ = s
	}
}

func BenchmarkSliceAllocZeroTellRequired(b *testing.B) {
	for b.Loop() {
		s := make([]int, 0, 10)
		for range 10 {
			s = append(s, 1)
		}

		_ = s
	}
}
