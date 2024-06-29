package smi

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	testShapes := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{3.0, 6.5}, 3.0 * 6.5},
		{Circle{4.5}, 63.61725123519331},
	}

	for _, tc := range testShapes {
		got := tc.shape.Area()
		want := tc.want

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	testShapes := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{Height: 3.0, Width: 6.5}, 2 * (3.0 + 6.5)},
		{"circle", Circle{Radius: 4.5}, 2 * math.Pi * 4.5},
	}

	for _, tc := range testShapes {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Perimeter()
			want := tc.want

			if got != want {
				t.Errorf("%#v got %g, want %g", tc.shape, got, want)
			}
		})
	}
}
