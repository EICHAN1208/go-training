package person

import "testing"

func TestBirthday(t *testing.T) {
	p := Person{Name: "Taro", Age: 20}
	p.Birthday()
	got := p.Age
	want := 21

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
