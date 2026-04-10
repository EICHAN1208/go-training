package max

import (
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "normal case", nums: []int{4, 9, 2, 7}, want: 9},
		{name: "empty case", nums: []int{}, want: 0},
		{name: "negative case", nums: []int{-4, -9, -2, -7}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.nums)
			want := tt.want
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
