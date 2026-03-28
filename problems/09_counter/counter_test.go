package counter

import "testing"

func TestNewCounter(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"once", 1},
		{"twice", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := NewCounter()
			var got int
			want := tt.want
			// wantの数だけcounter()を呼ぶ
			for i := 0; i < want; i++ {
				got = counter()
			}
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
