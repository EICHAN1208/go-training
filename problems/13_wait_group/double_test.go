package main

import (
	"reflect"
	"testing"
)

func TestDoubleAll(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "通常ケース",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "順序保持",
			nums: []int{5, 3, 1, 4, 2},
			want: []int{10, 6, 2, 8, 4},
		},
		{
			name: "空スライス",
			nums: []int{},
			want: []int{},
		},
		{
			name: "1要素",
			nums: []int{3},
			want: []int{6},
		},
		{
			name: "負の値",
			nums: []int{-1, -2},
			want: []int{-2, -4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DoubleAll(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleAll(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
