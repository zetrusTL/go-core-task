package main

import (
	"fmt"
	"sync"
	"time"
)

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 13; i++ {
			ch2 <- i
			time.Sleep(150 * time.Millisecond)
		}
		close(ch2)
	}()

	go func() {
		for i := 100; i < 103; i++ {
			ch3 <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch3)
	}()

	out := merge(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}
