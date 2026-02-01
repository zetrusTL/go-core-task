package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generation(ch chan<- int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer close(ch)
	ch <- r.Intn(1000000)
	time.Sleep(time.Millisecond * 500)
}

func main() {
	ch := make(chan int)
	res := 0
	go generation(ch)
	res = <-ch
	fmt.Println(res)
}