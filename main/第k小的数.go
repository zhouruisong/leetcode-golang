package main

import "fmt"

//第k小的数
func findKthMin(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	//转换成求第len(nums)-k+1大的数
	return findKthLargest(nums, len(nums)-k+1)
}

func findKthLargest(nums []int, k int) int {
	//快排序思想
	return partition(nums, k)
}

func partition(nums []int, k int) int {
	i := 0
	j := len(nums) - 1
	loc := nums[i]
	for i < j {
		for i < j && nums[j] >= loc {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

		for i < j && nums[i] <= loc {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i] = loc

	if len(nums)-i == k {
		return nums[i]
	}
	if len(nums)-i > k {
		return partition(nums[i+1:], k)
	}
	return partition(nums[:i], k+i-len(nums))
}

func main() {
	arr := []int{9, 1, 3, 2, 4, 6}
	fmt.Println(findKthMin(arr, 3))
}
