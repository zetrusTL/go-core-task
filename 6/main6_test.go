package main

import (
	"testing"
	"time"
)

func TestGenerationSendsValueAndCloses(t *testing.T) {
	ch := make(chan int)

	go generation(ch)

	select {
	case v, ok := <-ch:
		if !ok {
			t.Fatal("channel closed before value received")
		}
		if v < 0 || v >= 1000000 {
			t.Fatalf("value out of range: %d", v)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("timeout waiting for value")
	}

	select {
	case _, ok := <-ch:
		if ok {
			t.Fatal("expected channel to be closed after sending one value")
		}
	case <-time.After(800 * time.Millisecond):
		t.Fatal("timeout waiting for channel to close")
	}
}

func TestGenerationDoesNotPanicOnCloseBySender(t *testing.T) {
	ch := make(chan int)
	go generation(ch)

	_, _ = <-ch
	_, ok := <-ch
	if ok {
		t.Fatal("expected channel closed")
	}
}
