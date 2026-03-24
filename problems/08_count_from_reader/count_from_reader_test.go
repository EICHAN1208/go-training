package count_from_reader

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountFromReader(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{"empty", "", map[string]int{}},
		{"alphabets only", "abc", map[string]int{"a": 1, "b": 1, "c": 1}},
		{"with non-alphabets", "a1b!c", map[string]int{"a": 1, "b": 1, "c": 1}},
		{"upper and lower", "aAbB", map[string]int{"a": 1, "A": 1, "b": 1, "B": 1}},
		{"repeated chars", "banana", map[string]int{"b": 1, "a": 3, "n": 2}},
		{"across buffer boundary", "abcdefghijk", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1, "g": 1, "h": 1, "i": 1, "j": 1, "k": 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountFromReader(strings.NewReader(test.input))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
