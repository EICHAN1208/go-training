package shape

import (
	"math"
	"testing"
)

func TestTotalArea(t *testing.T) {
	tests := []struct {
		name   string
		shapes []Shape
		want   float64
	}{
		{
			name:   "only rectangle",
			shapes: []Shape{Rectangle{Width: 4, Height: 5}},
			want:   20,
		},
		{
			name:   "only circle",
			shapes: []Shape{Circle{Radius: 3}},
			want:   math.Pi * 9,
		},
		{
			name:   "circle and rectangle",
			shapes: []Shape{Circle{Radius: 3}, Rectangle{Width: 4, Height: 5}},
			want:   math.Pi*9 + 20,
		},
		{
			name:   "empty",
			shapes: []Shape{},
			want:   0,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := TotalArea(v.shapes)

			// 浮動小数点の誤差対策
			if math.Abs(got-v.want) > 1e-9 {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
