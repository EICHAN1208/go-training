package authstore

import (
	"errors"
	"strings"
	"testing"
)

func TestCallAPI_OK(t *testing.T) {
	v, err := CallAPI("token", 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != "ok" {
		t.Errorf("got %q, want %q", v, "ok")
	}
}

func TestCallAPI_Unauthorized(t *testing.T) {
	_, err := CallAPI("", 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.HasPrefix(err.Error(), "CallAPI: ") {
		t.Errorf("error should be wrapped with 'CallAPI: ' prefix, got: %v", err)
	}

	// errors.Is で ErrUnauthorized が判定できること（%w でラップしている必要あり）
	if !errors.Is(err, ErrUnauthorized) {
		t.Errorf("errors.Is(err, ErrUnauthorized) should be true")
	}
}

func TestCallAPI_QuotaExceeded(t *testing.T) {
	_, err := CallAPI("token", 120)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.HasPrefix(err.Error(), "CallAPI: ") {
		t.Errorf("error should be wrapped, got: %v", err)
	}

	// errors.As で *QuotaExceededError を取り出せること
	var qe *QuotaExceededError
	if !errors.As(err, &qe) {
		t.Fatalf("errors.As should extract *QuotaExceededError, got: %v", err)
	}
	if qe.Limit != 100 {
		t.Errorf("Limit: got %d, want %d", qe.Limit, 100)
	}
	if qe.Used != 120 {
		t.Errorf("Used: got %d, want %d", qe.Used, 120)
	}
}

func TestClassifyError(t *testing.T) {
	tests := []struct {
		name  string
		token string
		quota int
		want  string
	}{
		{name: "ok", token: "token", quota: 50, want: "ok"},
		{name: "unauthorized", token: "", quota: 0, want: "unauthorized"},
		{name: "quota_exceeded", token: "token", quota: 120, want: "quota_exceeded:120/100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CallAPI(tt.token, tt.quota)
			got := ClassifyError(err)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClassifyError_Unknown(t *testing.T) {
	// 既知のエラーに該当しない場合は "unknown"
	got := ClassifyError(errors.New("something else"))
	if got != "unknown" {
		t.Errorf("got %q, want %q", got, "unknown")
	}
}

// errors.Is と errors.As の使い分けに関する説明テスト。
//
//	❌ よくある間違い:
//	  if _, ok := err.(*QuotaExceededError); ok { ... }
//	  → ラップされていると false になる（型アサーションは1段しか見ない）
//
//	✅ 正解:
//	  var qe *QuotaExceededError
//	  if errors.As(err, &qe) { ... }
//	  → エラーチェーンを辿って取り出してくれる
//
//	センチネルエラーは errors.Is、構造体エラーは errors.As。
func TestClassifyError_DocReminder(t *testing.T) {
	t.Log("errors.Is = 値での識別 / errors.As = 型での識別 + フィールド取り出し")
}
