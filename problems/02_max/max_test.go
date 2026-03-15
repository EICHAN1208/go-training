package max

import (
	"testing"
)

func TestMax(t *testing.T) {
	max := Max([]int{4, 9, 2, 7})
	want := 9
	if max != want {
		t.Errorf("got %d, want %d", max, want)
	}
}
