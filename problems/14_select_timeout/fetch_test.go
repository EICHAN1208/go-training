package fetch

import (
	"testing"
	"time"
)

func TestFetchTimeout(t *testing.T) {
	tests := []struct {
		name    string
		delay   time.Duration
		timeout time.Duration
		want    string
		wantErr bool
	}{
		{
			name:    "normal case",
			delay:   100 * time.Millisecond,
			timeout: 200 * time.Millisecond,
			want:    "result",
			wantErr: false,
		},
		{
			name:    "timeout case",
			delay:   300 * time.Millisecond,
			timeout: 100 * time.Millisecond,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchWithTimeout(tt.delay, tt.timeout)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, but got nil")
				} else if err.Error() != "timeout" {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tt.want {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
