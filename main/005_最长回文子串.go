package main

import "fmt"

/*
* 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
*/

func isPalindromeStr(s string) bool {
	//fmt.Println(s)
	n := len(s)
	j := n - 1
	i := 0
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	j := 0

	var max string
	for j <= n {
		i := j
		for i < n {
			//是否是回文
			//fmt.Println(s[j : i+1])
			if isPalindromeStr(s[j : i+1]) {
				if len(s[j:i+1]) > len(max) {
					//fmt.Println(s[j : i+1])
					max = s[j : i+1]
					//fmt.Println(j,i+1)
				}
			}
			i++
		}
		j++
	}

	fmt.Println(max)
	return max
}

func main() {
	//fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("babad"))
}
