package sum

import "testing"

func Test_Sum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 5})
	want := 15
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
