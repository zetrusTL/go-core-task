package main

import(
	"fmt"
)

func Diff(slice1 []string, slice2 []string) []string {
	m := make(map[string]bool)
	for _, val := range slice2{
		m[val] = true
	}
	res := make([]string,0,(len(slice1)))

	for _,val := range slice1{
		if !m[val]{
			res = append(res,val)
		}
	}
	return res
}

func main(){
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Diff(slice1,slice2))
}