package main

import "fmt"

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (46.88%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    116.8K
 * Total Submissions: 249.1K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 *
 * 说明:
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 */

func mergea(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	total := m + n
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[total-1] = nums1[i]
			total--
			i--
		} else {
			nums1[total-1] = nums2[j]
			total--
			j--
		}
	}

	if i < 0 && j >= 0 {
		for j >= 0 {
			nums1[total-1] = nums2[j]
			total--
			j--
		}
	}

	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergea(nums1, 4, nums2, 3)
}
