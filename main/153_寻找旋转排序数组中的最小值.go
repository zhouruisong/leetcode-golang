package main

import "fmt"

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
*/

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left := 0
	right := n - 1
	for left < n && nums[left] > nums[right] {
		left++
	}

	return nums[left]
}

func findMin2(nums []int) int {
	lens := len(nums)
	left := 0
	right := lens - 1
	mid := right / 2

	if left == mid {
		if nums[left] < nums[right] {
			return nums[left]
		}
		return nums[right]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}

	for {
		if nums[left] <= nums[mid] {
			left = mid
			mid = (left + right) / 2
		} else if nums[left] > nums[mid] {
			right = mid
			mid = (left + right) / 2
		}

		if mid == left {
			return nums[mid+1]
		}
	}
}

func main() {
	x := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(x))
}
