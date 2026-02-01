package main

import "testing"

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	out := merge(ch1, ch2)

	sum := 0
	count := 0

	for v := range out {
		sum += v
		count++
	}

	if count != 4 {
		t.Fatal("expected 4 values, got", count)
	}

	if sum != 10 {
		t.Fatal("unexpected sum:", sum)
	}
}
