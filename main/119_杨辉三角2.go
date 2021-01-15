package main

import "fmt"

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (62.38%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    83.4K
 * Total Submissions: 133.3K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

func getRow(rowIndex int) []int {
	//dp算法
	dp := make([][]int, rowIndex+1)
	if rowIndex == 0 {
		return []int{1}
	}

	for k := range dp {
		dp[k] = make([]int, k+1)
	}

	for i := 0; i < rowIndex+1; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}

	return dp[rowIndex]
}

func main() {
	fmt.Print(getRow(3))
}
