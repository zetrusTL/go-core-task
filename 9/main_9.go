package main

import "fmt"

func Pipeline(in <-chan uint8) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for v := range in {
			f := float64(v)
			out <- f * f * f
		}
	}()

	return out
}

func main() {
	in := make(chan uint8)
	out := Pipeline(in)

	go func() {
		defer close(in)
		for _, v := range []uint8{1, 2, 3, 4, 5} {
			in <- v
		}
	}()

	for v := range out {
		fmt.Println(v)
	}
}
