package main

import "fmt"

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (69.66%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    145.4K
 * Total Submissions: 208.3K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */

func generate(numRows int) [][]int {
	//dp算法
	dp := make([][]int, numRows)
	if numRows == 0 {
		return dp
	}

	for k := range dp {
		dp[k] = make([]int, k+1)
	}

	for i:=0;i<numRows;i++{
		for j:=0;j<i+1;j++{
			if j==0 || j==i {
				dp[i][j] = 1
			}else{
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}

	return dp
}

func main()  {
	fmt.Print(generate(5))
}
