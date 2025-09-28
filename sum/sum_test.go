package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertSum(t, got, want, numbers)
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertSum(t, got, want, numbers)
	})
}

func assertSum(t *testing.T, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
	}
}
