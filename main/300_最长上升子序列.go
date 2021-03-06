package main

import (
	"fmt"
)

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

//O(n^2)
func lengthOfLIS(nums []int) int {
	/*
		    [1,3,2,3,1,4]
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

	//O(n^2)
	//[1,3,2,3,1,4]
	n := len(nums)
	if n < 2 {
		return n
	}

	//dp算法
	//dp[i]表示以nums[i]结尾的最长递增子序列长度
	dp := make([]int, n)
	dp[0] = 1

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	//dp[i]表示以i下标数值为子序列结束时最长上升子序列的长度,初始时默认为1
	// for i := 0; i < len(dp); i++ {
	// 	dp[i] = 1
	// }

	ret := 0
	for i := 1; i < n; i++ {
		dp[i] = 1 //默认值
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if ret < dp[i] {
			ret = dp[i]
		}
	}

	return ret
}

/**
 * retrun the longest increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型一维数组
 */
//o(nlogn)
func LIS(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	//记录以i结尾的最大上升子序列长度
	maxLen := make([]int, len(arr))
	//记录最大上升子序列
	res := []int{arr[0]}
	maxLen[0] = 1
	for i := 1; i < len(arr); i++ {
		//当前值比res的最后一个值大
		if arr[i] > res[len(res)-1] {
			res = append(res, arr[i])
			maxLen[i] = len(res)
		} else {
			//当前值比res的最后一个值小，需要找到第一个大于等于当前值的位置, 该位置的值被当前值替换
			l, r := 0, len(res)-1
			for l < r {
				mid := l + (r-l)>>1
				if res[mid] >= arr[i] {
					r = mid
				} else {
					l = mid + 1
				}
			}
			res[l] = arr[i]
			maxLen[i] = l + 1
		}
	}
	//从maxLen后往前找，找到被替换过的位置的值
	for i, j := len(maxLen)-1, len(res); j > 0; i-- {
		if maxLen[i] == j {
			j--
			res[j] = arr[i]
		}
	}
	return res
}

func main() {
	//arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//arr := []int{1, 3, 2, 3, 1, 4}
	arr := []int{1, 2, 8, 6, 4}
	//fmt.Println(lengthOfLIS(arr))
	fmt.Println(LIS(arr))
}
