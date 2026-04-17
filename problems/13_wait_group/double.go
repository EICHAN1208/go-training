package main

import (
	"sync"
)

type result struct{ index, value int }

func DoubleAll(nums []int) []int {
	s := make([]int, len(nums))
	ch := make(chan result)
	var wg sync.WaitGroup

	// 各numsの要素を別々のgoroutineで処理
	for i, v := range nums {
		wg.Add(1) // goroutine起動前に+1する

		// goroutine起動
		go func(v, i int) {
			defer wg.Done() // 関数を抜けるときに必ず-1する
			ch <- result{index: i, value: v * 2}
		}(v, i)
	}

	// ここでwg.Wait()とclose(ch)を別のgoroutineで実行する理由
	// main goroutineが wg.Wait() でブロックされると ch が受信できないためデッドロックが発生する
	go func() {
		wg.Wait() // カウンターが0になるまでここでブロックする
		close(ch) // すべてのgroutineが Done したらchannelをクローズする
	}()

	// 各要素が2倍になったスライスを返す
	for r := range ch {
		s[r.index] = r.value
	}
	return s
}
