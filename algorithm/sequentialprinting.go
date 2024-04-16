package main

import (
	"fmt"
	"sync"
)

// 3 个协程，每个协程只打印 a/b/c 中的一个，要求按照 a/b/c 的顺序循环打印 100 次，示例：
// a
// b
// c
// a
// b
// c

var wg = &sync.WaitGroup{}

func printChar(alphabet string, times int, seqChan []chan struct{}, seq int) {
	wg.Add(1)

	for i := 0; i < times; i++ {
		select {
		case <-seqChan[seq]:
			fmt.Println(alphabet)
		}

		nextSeq := (seq + 1) % len(seqChan)
		seqChan[nextSeq] <- struct{}{}
	}

	if seq == 0 {
		<-seqChan[seq]
	}

	wg.Done()
}

func main() {
	alphabets := []string{"a", "b", "c"}
	printTimes := 10

	seqChan := make([]chan struct{}, len(alphabets))

	for i := range seqChan {
		seqChan[i] = make(chan struct{})
	}

	for i, c := range alphabets {
		go printChar(c, printTimes, seqChan, i)
	}

	seqChan[0] <- struct{}{}

	wg.Wait()
}
