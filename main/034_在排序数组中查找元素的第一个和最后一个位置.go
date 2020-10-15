package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1,-1}
	}

	left := 0
	right := n - 1
	mid := 0
	//二分查找思想
	for left <= right {
		mid = (right + left) >> 1
		if nums[mid] == target {
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(mid)
	if nums[mid] != target {
		return []int{-1,-1}
	}

	ret := make([]int,2)
	//从当前位置向左遍历，获取到开始节点
	for i := mid; i >= 0; i-- {
		if nums[i] == target {
			fmt.Println(i)
			ret[0] = i
		} else {
			break
		}
	}

	//从当前位置向右遍历，获取结束节点
	for i := mid; i < n; i++ {
		if nums[i] == target {
			fmt.Println(i)
			ret[1] = i
		} else {
			break
		}
	}

	return ret
}

func main() {
	fmt.Println(searchRange([]int{1,4}, 4))
}
