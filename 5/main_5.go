package main
import(
	"fmt"
)

func Diff(num1, num2 []int) (bool,[]int) {
	m:=make(map[int]bool)

	for _,val := range num1{
		m[val] = true
	}

	res := make([]int,0,len(num1))
	flag := false
	for _,val := range num2{
		if m[val] {
			flag = true
			res = append(res, val)
		}
	}

	return flag, res
}

func main(){
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Diff(a,b))
}