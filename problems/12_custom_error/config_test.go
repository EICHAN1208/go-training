package config

import (
	"errors"
	"testing"
)

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantVal  string
		wantErr  *FileError
	}{
		{name: "empty", filename: "", wantErr: &FileError{FileName: "", Reason: "filename is empty"}},
		{name: "missing text file", filename: "missing.txt", wantErr: &FileError{FileName: "missing.txt", Reason: "file not found"}},
		{name: "ok", filename: "config.json", wantVal: "debug=true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := ReadConfig(tt.filename)

			if err == nil && tt.wantErr == nil {
				if v != tt.wantVal {
					t.Errorf("got %v, want %v", v, tt.wantVal)
				}
			} else if err == nil && tt.wantErr != nil {
				t.Errorf("expected error but got nil")
			} else {
				var fe *FileError
				if errors.As(err, &fe) {
					// ファイル名が違っていたらエラー出力する
					if fe.FileName != tt.wantErr.FileName {
						t.Errorf("got %v, want %v", fe.FileName, tt.wantErr.FileName)
					}

					// エラー理由が違っていたらエラー出力する
					if fe.Reason != tt.wantErr.Reason {
						t.Errorf("got %v, want %v", fe.Reason, tt.wantErr.Reason)
					}
				} else {
					t.Errorf("unexpected error type: %v", err)
				}
			}
		})
	}
}
