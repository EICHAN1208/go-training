package count_letters

import (
	"reflect"
	"testing"
)

func TestCountLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{"banana", "banana", map[string]int{"b": 1, "a": 3, "n": 2}},
		{"empty", "", map[string]int{}},
		{"single", "a", map[string]int{"a": 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountLetters(test.input)
			want := test.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
