package main

import "fmt"

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输出: false
*/

//https://leetcode-cn.com/problems/wildcard-matching/solution/yi-ge-qi-pan-kan-dong-dong-tai-gui-hua-dpsi-lu-by-/

func isMatch(s string, p string) bool {
	heng := len(s) + 1
	zong := len(p) + 1

	dp := make([][]bool, heng)
	for i := range dp {
		dp[i] = make([]bool, zong)
	}

	dp[0][0] = true

	if p[0] == '*' {
		for i := 1; i < zong; i++ {
			dp[1][i] = true
		}
	}

	for i := 1; i < heng; i++ {
		path := false
		for j := 1; j < zong; j++ {
			if p[i-1] == '*' {
				if dp[i-1][0] {
					for k := 1; k < zong; k++ {
						dp[i][k] = true
					}
				}

				if dp[i-1][j] {
					path = true
				}
				if path {
					dp[i][j] = true
				}
			} else if p[i-1] == '?' || p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[heng-1][zong-1]
}

func main() {
	//s := "aa"
	//p := "a"

	s := "aa"
	p := "a"

	fmt.Println(isMatch(s, p))
}

//class Solution:
//	def isMatch(self, s: str, p: str) -> bool:
//	#s横，p纵
//	#土味拼音变量名
//	zong = len(p)+1 #纵轴长度
//	heng = len(s)+1 #横轴长度
//
//	table = [[False]*heng for i in range(zong)]
//	table[0][0] = True
//
//	if p.startswith("*"):
//		table[1] = [True]*heng
//
//	for m in range(1,zong):
//	    path = False #是否可以在该行使用*的特技
//		for n in range(1,heng):
//			if p[m-1] == "*": #m是表格的索引，但不是p的索引
//				if table[m-1][0] == True:
//					table[m] = [True]*heng
//				if table[m-1][n]: #只要顶上有了True，就可以开通*接下来的所有道路
//					path = True
//				if path:
//					table[m][n] = True
//			elif p[m-1] == "?" or p[m-1]==s[n-1]: #先判断字母是否符合
//				table[m][n] = table[m-1][n-1] #再看左上方格子是不是T
//
//	return table[zong-1][heng-1]
