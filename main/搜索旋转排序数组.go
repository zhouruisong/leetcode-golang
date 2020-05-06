package main

import "fmt"

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

/*
这是一种梳轴问题的通用解法。
算法思路：

有mid和left处的值判断mid那边是单调递增区间
对于递增区间，由区间端点的值，可以判断target是否肯定(注意这个定语)在此区间内，然后调整left和right；否则，说明在相反的区间内。
*/
func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if nums[mid] > target && nums[left] <= target { //1.1 [left,mid]这个区间,通过端点可以肯定target是否在此区间
				right = mid - 1
			} else { //1.2 不在，说明在[mid,right]区间
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	//判断目标值与头节点的大小，如果大则必定在前半段，小则必定是在后半段
	if target == nums[0] {
		return 0
	}
	l := len(nums)
	if target == nums[l-1] {
		return l - 1
	}
	if target > nums[0] {
		i := 0
		for i+1 < len(nums)-1 && nums[i] < nums[i+1] {
			if target == nums[i+1] {
				return i + 1
			}
			i++
		}
		return -1
	}

	if target < nums[0] {
		i := l - 1
		for i-1 >= 0 && nums[i-1] < nums[i] {
			if target == nums[i-1] {
				return i - 1
			}
			i--
		}
		return -1
	}

	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(arr, 2))
}
