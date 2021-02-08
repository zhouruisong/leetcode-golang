package main

import "fmt"

//最长公共子串要求在原字符串中是连续的，而子序列只需要保持相对顺序一致，并不要求连续。
func longestCommonSubstr(text1 string, text2 string) int {
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

	max := -1
	index := 0
	//str := []string{}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			//不相同， 为0
			if text1[i-1] != text2[j-1] {
				dp[i][j] = 0
			} else { //如果末端相同，等于前面最长公共子串数+1
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if max < dp[i][j] {
				max = dp[i][j]
				index = i //记录下最长子串的末尾字符的位置
			}
		}
	}

	//fmt.Println(max)
	//fmt.Println(index)
	//
	fmt.Println(text1[index-max : index])
	return max
}

func LCS(str1 string, str2 string) string {
	// write code here
	n1 := len(str1)
	n2 := len(str2)

	if n1 == 0 && n2 == 0 {
		return "-1"
	}

	if n1 == 0 || n2 == 0 {
		return "-1"
	}

	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	max := 0
	index := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			//不相同， 为0
			if str1[i-1] != str2[j-1] {
				dp[i][j] = 0
			} else { //如果末端相同，等于前面最长公共子串数+1
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if max < dp[i][j] {
				max = dp[i][j]
				index = i
			}
		}
	}

	//fmt.Println(dp)
	if max == 0 {
		return "-1"
	}

	return str1[index-max : index]
}

func main() {
	//s1 := "abbecde"
	//s2 := "abe"

	s1 := "1AB2345CD"
	s2 := "12345EF"
	fmt.Println(longestCommonSubstr(s1, s2))
	fmt.Println(LCS(s1, s2))
}
