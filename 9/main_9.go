package main
import (
	"fmt"
)

func Pipeline(ch <-chan uint8) (<- chan float64){
	ch1 := make(chan float64)
	go func() {
		defer close(ch1)
		for val := range ch{
			f := float64(val)
			ch1 <- f*f*f
		}
	}()

	return ch1
}

func main() {
	ch := make(chan uint8)
	res:= Pipeline(ch)

	go func(){
		defer close(ch)
		for _,val:= range []uint8{1,2,3,4,5}{
			ch <- val
		}
	}()

		for v := range res {
		fmt.Println(v)
	}
}