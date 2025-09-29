package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertFloat(t, want, got)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0
	assertFloat(t, want, got)
}

func assertFloat(t testing.TB, want, got float64) {
	t.Helper()

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}
