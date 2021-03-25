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
//https://leetcode-cn.com/problems/edit-distance/solution/bian-ji-ju-chi-mian-shi-ti-xiang-jie-by-labuladong/
//o(mn)
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

	//我们用 D[i][j] 表示 A 的前 i 个字母和 B 的前 j 个字母之间的编辑距离
	//如上所述，当我们获得 D[i][j-1]，D[i-1][j] 和 D[i-1][j-1] 的值之后就可以计算出 D[i][j]。
	//
	//D[i][j-1] 为 A 的前 i 个字符和 B 的前 j - 1 个字符编辑距离的子问题。
	//即对于 B 的第 j 个字符，我们在 A 的末尾添加了一个相同的字符，那么 D[i][j] 最小可以为 D[i][j-1] + 1；

	//D[i-1][j] 为 A 的前 i - 1 个字符和 B 的前 j 个字符编辑距离的子问题。
	//即对于 A 的第 i 个字符，我们在 B 的末尾添加了一个相同的字符，那么 D[i][j] 最小可以为 D[i-1][j] + 1；

	//D[i-1][j-1] 为 A 前 i - 1 个字符和 B 的前 j - 1 个字符编辑距离的子问题。
	//即对于 B 的第 j 个字符，我们修改 A 的第 i 个字符使它们相同，那么 D[i][j] 最小可以为 D[i-1][j-1] + 1。
	//特别地，如果 A 的第 i 个字符和 B 的第 j 个字符原本就相同，那么我们实际上不需要进行修改操作。
	//在这种情况下，D[i][j] 最小可以为 D[i-1][j-1]。
	//对于边界情况，一个空串和一个非空串的编辑距离为 D[i][0] = i 和 D[0][j] = j，D[i][0]
	//相当于对 word1 执行 i 次删除操作，D[0][j] 相当于对 word1执行 j 次插入操作。

	/*
		那么我们可以写出如下的状态转移方程：
		若 A 和 B 的最后一个字母相同：
		D[i][j]=min(D[i][j−1]+1,D[i−1][j]+1,D[i−1][j−1])
		若 A 和 B 的最后一个字母不相同：
		D[i][j]=min(D[i][j−1]+1,D[i−1][j]+1,D[i−1][j−1])

		时间复杂度 ：O(mn)O(mn)，其中 mm 为 word1 的长度，nn 为 word2 的长度。
		空间复杂度 ：O(mn)O(mn)，我们需要大小为 O(mn)O(mn) 的 DD 数组来记录状态值。
	*/
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//最后两个字符相同
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				//fmt.Println("啥都不做", i, j)
			} else {
				dp[i][j] = min3(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}

	//for i := 0; i <= m; i++ {
	//	fmt.Println(dp[i])
	//}
	// 储存着整个 s1 和 s2 的最小编辑距离
	return dp[m][n]
}

func min3(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
