package main

import "fmt"

/*
* 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
 *
 * 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde"
 * 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
 *
 * 若这两个字符串没有公共子序列，则返回 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入：text1 = "abcde", text2 = "ace"
 * 输出：3
 * 解释：最长公共子序列是 "ace"，它的长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc"，它的长度为 3。
 *
 *
 * 示例 3:
 *
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0。
*/

/*
解法
例如：s1="abcde"　　s2= "ace"，求两个字符串的公共子序列，答案是“ace”

1.　S{s1,s2,s3....si} T{t1,t2,t3,t4....tj}

2.　子问题划分

(1) 如果S的最后一位等于T的最后一位，则最大子序列就是{s1,s2,s3...si-1}和{t1,t2,t3...tj-1}的最大子序列+1

(2) 如果S的最后一位不等于T的最后一位，那么最大子序列就是

① {s1,s2,s3..si}和 {t1,t2,t3...tj-1} 最大子序列

② {s1,s2,s3...si-1}和{t1,t2,t3....tj} 最大子序列

以上两个自序列的最大值

3.　边界

只剩下{s1}和{t1}，如果相等就返回1，不等就返回0

4.　使用一个表格来存储dp的结果

如果 S[i] == T[j] 则dp[i][j] = dp[i-1][j-1] + 1

否则dp[i][j] = max(dp[i][j-1],dp[i-1][j])
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)

	if n1 == 0 && n2 == 0 {
		return 0
	}

	if n1 == 0 || n2 == 0 {
		return 0
	}

	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	fmt.Println(dp)

	str := []string{}
	max := -1
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			//如果末端相同
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

			} else { //如果末端不相同
				dp[i][j] = findmax(dp[i-1][j], dp[i][j-1])
			}

			//记录最大值，其中i
			if max < dp[i][j] {
				max = dp[i][j]
				str = append(str, string(text1[i-1]))
				fmt.Println(i, j)
			}
		}
	}

	fmt.Println(str)
	return dp[n1][n2]
}

func findmax(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	s1 := "abcd1e"
	s2 := "acde"
	fmt.Println(longestCommonSubsequence(s1, s2))
}
