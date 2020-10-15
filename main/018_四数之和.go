package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 4 {
		return result
	}
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}

		target := target - nums[i]
		ret := threeSum2(nums[i+1:], target)
		for j := 0; j < len(ret); j++ {
			result = append(result, append(ret[j], nums[i]))
		}
	}

	return result
}

//复杂度o(n*n),采用二分查找方法
func threeSum2(nums []int, target int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}

		L := i + 1
		R := n - 1

		for L < R {
			if nums[i]+nums[L]+nums[R] == target {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				//去重
				for L < R && nums[L] == nums[L+1] {
					L = L + 1
				}
				//去重
				for L < R && nums[R] == nums[R-1] {
					R = R - 1
				}

				L = L + 1
				R = R - 1
			} else if nums[i]+nums[L]+nums[R] < target {
				L = L + 1
			} else {
				R = R - 1
			}
		}
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
