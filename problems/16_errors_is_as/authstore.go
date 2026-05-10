package authstore

import (
	"errors"
	"fmt"
)

// ErrUnauthorized はセンチネルエラー（値で識別するエラー）。
// 呼び出し側は errors.Is(err, ErrUnauthorized) で判定する。
var ErrUnauthorized = errors.New("unauthorized")

// QuotaExceededError は構造体エラー（フィールドに情報を持つエラー）。
// 呼び出し側は errors.As で取り出してから Limit / Used を参照する。
type QuotaExceededError struct {
	Limit int
	Used  int
}

func (e *QuotaExceededError) Error() string {
	return fmt.Sprintf("quota exceeded: limit=%d, used=%d", e.Limit, e.Used)
}

// CallAPI は API 呼び出しを模した関数。
//
// 仕様:
//   - token == ""    のとき: ErrUnauthorized を fmt.Errorf("CallAPI: %w", err) でラップして返す
//   - quota >= 100   のとき: &QuotaExceededError{Limit: 100, Used: quota} を fmt.Errorf("CallAPI: %w", err) でラップして返す
//   - それ以外         のとき: "ok" と nil を返す
//
// チェック順は token → quota の順とする。
func CallAPI(token string, quota int) (string, error) {
	// TODO: 実装する
	return "", nil
}

// ClassifyError は err の中身を判定して文字列ラベルを返す。
//
// 仕様:
//   - err == nil                                            → "ok"
//   - errors.Is(err, ErrUnauthorized) が true                → "unauthorized"
//   - errors.As(err, &QuotaExceededError) が true            → "quota_exceeded:Used/Limit" の形式
//     例: Used=120, Limit=100 のとき "quota_exceeded:120/100"
//   - 上記以外                                                → "unknown"
//
// errors.Is と errors.As の使い分けがポイント。
//   - センチネルエラー（値で識別） → errors.Is
//   - 構造体エラー（型で識別、フィールドを使いたい） → errors.As
func ClassifyError(err error) string {
	// TODO: 実装する
	return ""
}
