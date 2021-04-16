package main

import "fmt"

/*
一个重复字符串是由两个相同的字符串首尾拼接而成，例如abcabc便是长度为6的一个重复字符串，而abcba则不存在重复字符串。
给定一个字符串，请编写一个函数，返回其最长的重复字符子串。
若不存在任何重复字符子串，则返回0。
示例1
输入
复制
"ababc"
返回值
复制
4
说明
abab为最长的重复字符子串，长度为4
示例2
输入
复制
"abcab"
返回值
复制
0
*/

func solve(a string) int {
	// write code here
	/*
	   可以将两个字符串想像成两个连续的滑动窗口，并假设这个滑动窗口最大是字符串长度的一半，通过比较两个窗口的内容是否相同，
	   不相同的时候不断从左向右平移，完了之后，还是不相同，这时候就将窗口的大小调小一点，直到找到一个相同的，
	   这个时候窗口的长度×2就是最大字符串的大小
	*/
	if len(a) == 0 {
		return 0
	}
	n := len(a)
	maxLen := n / 2
	for i := maxLen; i >= 1; i-- {
		for j := 0; j <= n-2*i; j++ {
			if check(a, j, i) {
				return 2 * i
			}
		}
	}
	return 0
}

func check(a string, start, n int) bool {
	for i := start; i < start+n; i++ {
		if a[i] != a[i+n] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solve("abcabc"))
	fmt.Println(solve("abcab"))
}
