package main

import (
	"fmt"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

func backtrace(start int, nums []int, res *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	//去重复
	m := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[start], nums[i] = nums[i], nums[start]
		backtrace(start+1, nums, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrace(0, nums, &res)
	return res
}

func main() {
	x := []int{1, 1, 5}
	fmt.Println(permuteUnique(x))
}
