package main

import "fmt"

/* 给定一个无序的整数数组，找到其中最长上升子序列的长度。
*
* 示例:
*
* 输入: [10,9,2,5,3,7,101,18]
* 输出: 4
* 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
*
* 说明:
*
*
* 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
* 你算法的时间复杂度应该为 O(n^2) 。
*
*
* 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*
 */

func lengthOfLIS(nums []int) int {
	/*
		第i个状态dp[i]代表以第i个元素结尾的最长上升子序列长度
		dp[i-1]代表以第i-1个元素结尾的最长上升子序列长度
		nums[i]一定是dp[i]所对应的最长上升子序列中最大的元素（因为在末尾）
		如：
		dp[0] = 1,[1]
		dp[1] = 2,[1,3]
		dp[2] = 2,[1,2]
		dp[3] = 3,[1,2,3]
		dp[4] = 1,[1]
		dp[5] = ?

		dp[5]对应的nums[5] = 4
		大于dp[0]对应的nums[0],则 [1]+[4] = [1,4]
		大于dp[1]对应的nums[1],则 [1,3]+[4] = [1,3,4]
		大于dp[2]对应的nums[2],则 [1,2]+[4] = [1,2,4]
		大于dp[3]对应的nums[3],则 [1,2,3]+[4] = [1,2,3,4]
		大于dp[4]对应的nums[4],则 [1]+[4] = [1,4]
		故最终dp[5] = 4

		取dp[0] --- dp[n-1]中的最大值
	*/

	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	ret := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}

	return ret
}

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(arr))
}
