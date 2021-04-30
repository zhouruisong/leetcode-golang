package main

/*题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例: 给定 nums = [2, 7, 11, 15], target = 9 因为 nums[0] + nums[1] = 2 + 7 = 9 所以返回 [0, 1]*/

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		other := target - v
		if j, ok := mp[other]; ok {
			return []int{j, k}
		} else {
			mp[v] = k
		}
	}

	return []int{}
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 22
	ret2 := twoSum(input, target)
	for _, v := range ret2 {
		fmt.Printf("ret2=%+v\n", v)
	}
}
