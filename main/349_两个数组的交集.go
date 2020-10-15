package main

import (
	"fmt"
	"sort"
)

/*
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。

 */

func intersection(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 == 0 && n2 == 0 {
		return []int{}
	}

	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	res := []int{}
	mp := make(map[int]bool, n1)
	for i := 0; i < n1; i++ {
		if mp[nums1[i]] == false {
			mp[nums1[i]] = true
		}
	}

	for j := 0; j < n2; j++ {
		//去重复
		if j >= 1 && nums2[j] == nums2[j-1] {
			continue
		}

		if mp[nums2[j]] {
			res = append(res, nums2[j])
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
