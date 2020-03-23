package main

import (
	"fmt"
)

/*
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

func permute(nums []int, k, ln int) {
	if ln-k <= 0 {
		return
	}

	fmt.Printf("%+v\n", nums)
	for i := k; i < ln-k; i++ {
		swap(&nums[k], &nums[i+1])
		fmt.Printf("%+v\n", nums)
		swap(&nums[i+1], &nums[k])
	}

	//permute(nums, k+1, ln)
}

func swap(x *int, y *int) {
	other := *x
	*x = *y
	*y = other
}

func permuteMain(nums []int) {
	ln := len(nums)

	for i := 0; i < ln; i++ {
		swap(&nums[0], &nums[i])
		fmt.Printf("=== %+v\n", nums)

		var news []int
		for k := 0; k < ln; k++ {
			news = append(news, nums[k])
		}

		fmt.Printf("news:%+v\n", news)
		permute(news, 1, ln)
		swap(&nums[i], &nums[0])
		//fmt.Printf("for=== %+v\n", nums)

	}
}

func main() {
	x := []int{1, 2, 3}
	permuteMain(x)
}
