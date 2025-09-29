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
			shape: Rectangle{
				Height: 12.0,
				Width:  6.0,
			},
			want: 72.0,
		},
		{
			shape: Circle{
				Radius: 10.0,
			},
			want: 314.1592653589793,
		},
		{
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			want: 36.0,
		},
	}

	for _, tt := range tests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("want %g but got %g", tt.want, got)
		}
	}
}
