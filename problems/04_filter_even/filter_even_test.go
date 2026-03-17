package filter_even

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"mixed odd and even", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"empty", []int{}, []int{}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"negative and zero", []int{-2, -1, 0, 1, 2}, []int{-2, 0, 2}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FilterEven(test.input)
			want := test.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
