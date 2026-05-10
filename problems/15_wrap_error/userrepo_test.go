package userrepo

import (
	"errors"
	"strings"
	"testing"
)

func TestGetUserName_Success(t *testing.T) {
	name, err := GetUserName("alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "Alice" {
		t.Errorf("got %q, want %q", name, "Alice")
	}
}

func TestGetUserName_WrapsNotFoundError(t *testing.T) {
	_, err := GetUserName("missing")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// 1) service 層でラップされたメッセージが付与されていること
	if !strings.HasPrefix(err.Error(), "GetUserName: ") {
		t.Errorf("error message should be wrapped with 'GetUserName: ' prefix, got: %v", err)
	}

	// 2) ラップされていても errors.As で *NotFoundError を取り出せること（%w を使う必然性）
	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Fatalf("errors.As should extract *NotFoundError from wrapped error, got: %v", err)
	}
	if nfe.Resource != "user" {
		t.Errorf("Resource: got %q, want %q", nfe.Resource, "user")
	}
	if nfe.ID != "missing" {
		t.Errorf("ID: got %q, want %q", nfe.ID, "missing")
	}
}

func TestGetUserName_WrapsValidationError(t *testing.T) {
	_, err := GetUserName("")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.HasPrefix(err.Error(), "GetUserName: ") {
		t.Errorf("error message should be wrapped, got: %v", err)
	}

	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("errors.As should extract *ValidationError, got: %v", err)
	}
	if ve.Field != "id" {
		t.Errorf("Field: got %q, want %q", ve.Field, "id")
	}
}

// %v でラップしてしまうと型情報が落ちることを確認するためのコメントテスト。
//
//	❌ 間違い: fmt.Errorf("GetUserName: %v", err)  → errors.As が false になる
//	✅ 正解  : fmt.Errorf("GetUserName: %w", err)  → errors.As が true になる
func TestGetUserName_WrappingPreservesType(t *testing.T) {
	_, err := GetUserName("missing")

	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Errorf(`%%w を使ってラップしていますか？ %%v でラップすると型情報が落ちて errors.As が false になります。`)
	}
}
