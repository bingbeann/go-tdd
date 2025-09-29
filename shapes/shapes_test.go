package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("want %g but got %g", want, got)
	}
}

func TestArea(t *testing.T) {
	assertArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("want %g but got %g", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		assertArea(t, Rectangle{12.0, 6.0}, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		assertArea(t, Circle{10.0}, 314.1592653589793)
	})
}
