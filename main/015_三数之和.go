package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func twoSum(input []int, target int) [][]int {
	has := make(map[int]bool)
	drop := make(map[int]bool)
	result := [][]int{}

	for i := 0; i < len(input); i++ {
		other := target - input[i]
		if has[other] {
			//去重
			if drop[other] == false {
				result = append(result, []int{input[i], other})
				drop[other] = true
			}
		} else {
			has[input[i]] = true
		}
	}

	return result
}

//复杂度o(n*n),借用之前求两数之和的方法
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n == 0 {
		return result
	}

	//系统排序
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] { //去重
			continue
		}
		other1 := 0 - nums[i]
		target := []int{}
		target = append(target, nums[i+1:]...)

		ret := twoSum(target, other1)
		if ret != nil {
			for _, t := range ret {
				t = append(t, nums[i])
				result = append(result, t)
			}
			//fmt.Printf("result=%v\n", result)
		}
	}

	return result
}

/*
先将数组进行排序
从左侧开始，选定一个值为 定值 ，右侧进行求解，获取与其相加为 00 的两个值
类似于快排，定义首和尾
首尾与 定值 相加
等于 00，记录这三个值
小于 00，首部右移
大于 00，尾部左移
定值右移，重复该步骤
*/
//复杂度o(n*n),采用二分查找方法
func threeSum2(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}

		L := i + 1
		R := n - 1

		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
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
			} else if nums[i]+nums[L]+nums[R] < 0 {
				L = L + 1
			} else {
				R = R - 1
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4}))
}
