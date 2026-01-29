package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(nums []int, n int) []int {
	res := make([]int, 0, len(nums)+1)
	res = append(res, nums...)
	res = append(res, n)
	return res
}

func copySlice(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

func removeElement(nums []int, i int) []int {
	if i < 0 || i >= len(nums) {
		return copySlice(nums)
	}
	res := make([]int, 0, len(nums)-1)
	res = append(res, nums[:i]...)
	res = append(res, nums[i+1:]...)
	return res
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = r.Intn(100)
	}

	fmt.Println("originalSlice:", originalSlice)
	fmt.Println("even:", sliceExample(originalSlice))
	fmt.Println("add 777:", addElements(originalSlice, 777))

	cp := copySlice(originalSlice)
	originalSlice[0] = -1
	fmt.Println("original modified:", originalSlice)
	fmt.Println("copy (unchanged):", cp)

	fmt.Println("remove index 3:", removeElement(originalSlice, 3))
}
