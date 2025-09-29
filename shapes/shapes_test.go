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
	tests := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{12.0, 6.0},
			want:  72.0,
		},
		{
			shape: Circle{10.0},
			want:  314.1592653589793,
		},
	}

	for _, tt := range tests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("want %g but got %g", tt.want, got)
		}
	}
}
