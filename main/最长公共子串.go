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
				//fmt.Println(i, j)
				//打印子串
				fmt.Println(string(text1[i-1]))
			}
		}
	}

	fmt.Println(dp)

	return max
}

func main() {
	s1 := "abbecde"
	s2 := "abe"
	fmt.Println(longestCommonSubstr(s1, s2))
}
