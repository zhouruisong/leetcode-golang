package main

import (
	"fmt"
)

/*
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/

type NodeDp struct {
	Val    int
	Choice int
	// 0 代表啥都不做
	// 1 代表插入
	// 2 代表删除
	// 3 代表替换
}

//https://leetcode-cn.com/problems/edit-distance/solution/bian-ji-ju-chi-mian-shi-ti-xiang-jie-by-labuladong/
func minDistance(word1 string, word2 string) int {
	//二维数组初始化(m+1)*(n+1)
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// 自底向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				//fmt.Println("啥都不做", i, j)
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			// 2. word1 增加一个字符
			dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			// 3. word1 删除 word1[i]字符
			dp[i][j] = min(dp[i][j], dp[i-1][j]+1)

		}
	}

	//for i := 0; i <= m; i++ {
	//	fmt.Println(dp[i])
	//}
	// 储存着整个 s1 和 s2 的最小编辑距离
	return dp[m][n]
}

func minDistance2(word1 string, word2 string) int {
	//二维数组初始化(m+1)*(n+1)
	m, n := len(word1), len(word2)
	dp := make([][]NodeDp, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]NodeDp, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0].Val = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j].Val = j
	}

	// 自底向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j].Val = dp[i-1][j-1].Val
			} else {
				dp[i][j].Val = dp[i-1][j-1].Val + 1
			}
			// 2. word1 增加一个字符
			dp[i][j].Val = min2(dp[i][j].Val, dp[i][j-1].Val+1, 2)
			// 3. word1 删除 word1[i]字符
			dp[i][j].Val = min2(dp[i][j].Val, dp[i-1][j].Val+1, 3)
		}
	}

	for i := 0; i <= m; i++ {
		fmt.Println(dp[i])
	}

	// 储存着整个 s1 和 s2 的最小编辑距离
	return dp[m][n].Val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min2(a, b, choice int) int {
	if a < b {
		printChoice(choice)
		fmt.Println(a)
		return a
	}

	printChoice(choice)
	fmt.Println(b)
	return b
}

func printChoice(choice int) {
	if choice == 1 {
		fmt.Println("插入")
	} else if choice == 2 {
		fmt.Println("替换")
	} else if choice == 3 {
		fmt.Println("删除")
	}
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
