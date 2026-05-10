package userrepo

import (
	"fmt"
)

// NotFoundError はリソースが見つからないことを表すカスタムエラー型。
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: id=%s", e.Resource, e.ID)
}

// ValidationError は入力検証エラーを表すカスタムエラー型。
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: field=%s, message=%s", e.Field, e.Message)
}

// FetchUser は repository 層を想定した関数。
//
// 仕様:
//   - id == ""        のとき: &ValidationError{Field: "id", Message: "id is empty"} を返す
//   - id == "missing" のとき: &NotFoundError{Resource: "user", ID: "missing"} を返す
//   - id == "alice"   のとき: "Alice" と nil を返す
//   - それ以外        のとき: "" と nil を返す（このテストでは未使用）
func FetchUser(id string) (string, error) {
	switch id {
	case "":
		return "", &ValidationError{Field: "id", Message: "id is empty"}
	case "missing":
		return "", &NotFoundError{Resource: "user", ID: "missing"}
	case "alice":
		return "Alice", nil
	default:
		return "", nil
	}
}

// GetUserName は service 層を想定した関数。内部で FetchUser を呼ぶ。
//
// 仕様:
//   - FetchUser がエラーを返した場合は、必ず fmt.Errorf("...: %w", err) でラップして返すこと
//     ラップのプレフィックスは "GetUserName: "
//     例: "GetUserName: user not found: id=missing"
//   - 正常系では FetchUser の戻り値をそのまま返す
//
// テストでは「ラップ後のエラーから errors.As で *NotFoundError / *ValidationError を取り出せる」ことを検証する。
// %v でラップしてしまうと型情報が落ちて errors.As が false になり、テストが落ちる。
func GetUserName(id string) (string, error) {
	name, err := FetchUser(id)
	if err != nil {
		return "", fmt.Errorf("GetUserName: %w", err)
	}
	return name, nil
}
