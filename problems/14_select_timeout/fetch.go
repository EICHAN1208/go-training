package fetch

import (
	"errors"
	"time"
)

func FetchWithTimeout(delay time.Duration, timeout time.Duration) (string, error) {
	// channelを用意する
	ch := make(chan string)
	// goroutine 内で time.Sleep(delay) の後に "result" という文字列を channel に送る
	go func(delay time.Duration) {
		time.Sleep(delay)
		ch <- "result"
	}(delay)
	// select を使って、結果が来たときとタイムアウトしたときを分岐する
	// - タイムアウトしたとき(time.Afterを使う)
	//   - errors.New("timeout")を返す
	// - タイムアウトしないとき
	//   - "result"文字列を返す
	select {
	case r := <-ch:
		return r, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}
