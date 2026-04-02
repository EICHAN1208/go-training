package main

import "fmt"

func sendVal(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendVal(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
