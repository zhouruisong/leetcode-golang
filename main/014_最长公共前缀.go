package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z
*/

func longestCommonPrefix(strs []string) string {
	//看清楚是前缀！
	n := len(strs)
	if n == 0 {
		return ""
	}

	if n == 1 {
		return strs[0]
	}

	var ret []uint8
	max := ""
	stop := false
	for j := 0; j < len(strs[0]); j++ {
		ans := strs[0][j]
		i := 1
		for ; i < n; i++ {
			if j >= len(strs[i]) {
				break
			}
			if ans != strs[i][j] {
				stop = true
				break
			}
		}

		if stop {
			break
		}

		if i >= n {
			ret = append(ret, ans)
			//fmt.Println(fmt.Sprintf("=========%s", ret))
		}
	}

	max = fmt.Sprintf("%s", ret)
	return max
	//fmt.Println(max)
}

func main() {
	strs := []string{"flabower", "flabow", "flabght"}
	//strs := []string{"dog", "racecar", "car"}
	//strs := []string{"aca","cba"}
	//strs := []string{"c","c"}
	fmt.Printf("ret=%v\n", longestCommonPrefix(strs))
}
