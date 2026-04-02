package main

import "testing"

func TestSendVal(t *testing.T) {
	ch := make(chan int)
	go sendVal(ch)

	expected := []int{1, 2, 3, 4, 5}
	for i, want := range expected {
		got, ok := <-ch
		if !ok {
			t.Fatalf("チャネルが早く閉じられた: %d個目で閉じた（5個期待）", i)
		}
		if got != want {
			t.Errorf("%d個目: got %d, want %d", i+1, got, want)
		}
	}

	_, ok := <-ch
	if ok {
		t.Error("5個送信後にチャネルが閉じられていない")
	}
}
