package divide

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
		wantErr  bool
	}{
		{
			name:     "normal case",
			dividend: 10,
			divisor:  2,
			want:     5,
			wantErr:  false,
		},
		{
			name:     "divide by zero",
			dividend: 10,
			divisor:  0,
			want:     0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.dividend, tt.divisor)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, but got nil")
				}
				if _, ok := err.(ErrZeroDivision); !ok {
					t.Errorf("unexpected error type: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
