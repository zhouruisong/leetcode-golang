package main

import "fmt"

/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

func strStr(haystack string, needle string) int {
	n2 := len(needle)
	if n2 == 0 {
		return 0
	}
	n1 := len(haystack)
	if n2 > n1 {
		return -1
	}

	loc := -1
	for i := 0; i < n1; i++ {
		j := 0
		k := i
		for ; j < n2; j++ {
			if needle[j] != haystack[k] {
				break
			}
			k++
		}

		if j >= n2 {
			loc = i
			break
		}
	}

	return loc
}

//KMP算法
func strStr2() {

}
func main() {
	haystack := "mississippi"
	needle := "mississippi"
	fmt.Println(strStr(haystack, needle))
}
