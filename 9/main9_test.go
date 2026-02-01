package main

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)

	out := Pipeline(in)

	go func() {
		defer close(in)
		in <- 2
		in <- 3
		in <- 4
	}()

	var res []float64
	for v := range out {
		res = append(res, v)
	}

	if len(res) != 3 {
		t.Fatal("expected 3 results, got", len(res))
	}

	if res[0] != 8 {
		t.Fatal("expected 8, got", res[0])
	}
	if res[1] != 27 {
		t.Fatal("expected 27, got", res[1])
	}
	if res[2] != 64 {
		t.Fatal("expected 64, got", res[2])
	}
}
