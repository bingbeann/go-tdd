package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertFloat(t, want, got)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0
	assertFloat(t, want, got)
}

func assertFloat(t testing.TB, want, got float64) {
	t.Helper()

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}
