package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000
*/
func merge(nums1 []int, nums2 []int) []int {
	result := []int{}
	n1 := len(nums1)
	n2 := len(nums2)
	l, r := 0, 0

	for l < n1 && r < n2 {
		if nums1[l] < nums2[r] {
			result = append(result, nums1[l])
			l++
		} else {
			result = append(result, nums2[r])
			r++
		}
	}

	result = append(result, nums1[l:]...)
	result = append(result, nums2[r:]...)
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//两个数组归并排序后，看长度
	arr := merge(nums1, nums2)

	fmt.Println(arr)
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}

	return float64(arr[n/2])
}

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2}

	fmt.Println(findMedianSortedArrays(arr1, arr2))
}
