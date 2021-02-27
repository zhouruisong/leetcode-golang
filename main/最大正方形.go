package main

/*
给定一个由0和1组成的2维矩阵，返回该矩阵中最大的由1组成的正方形的面积
示例1
输入
[[1,0,1,0,0],[1,0,1,1,1],[1,1,1,1,1],[1,0,0,1,0]]
返回值
4

二维动态规划，时间复杂度为O(m*n)。
用dp[i][j]表示以matrix[i][j]为右下角的全为1组成的最大正方形的边长，则:
1. 若matrix[i][j] == 0, 则dp[i][j] = 0
2. 若matrix[i][j] == 1, 则dp[i][j] = min{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]} + 1
*/

func solve(matrix [][]byte) int {
	// write code here
	n := len(matrix)
	m := len(matrix[0])
	ans := 0
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				}
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
