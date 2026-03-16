package count_letters

import (
	"reflect"
	"testing"
)

func TestCountLetters(t *testing.T) {
	got := CountLetters("banana")
	want := map[string]int{
		"a": 3,
		"b": 1,
		"n": 2,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
