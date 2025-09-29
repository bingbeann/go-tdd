package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assertFloat(t, want, got)
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		assertFloat(t, want, got)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		assertFloat(t, want, got)
	})
}

func assertFloat(t testing.TB, want, got float64) {
	t.Helper()

	if got != want {
		t.Errorf("want %g but got %g", want, got)
	}
}
