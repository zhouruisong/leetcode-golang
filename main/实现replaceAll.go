package main

import (
	"fmt"
	//"strings"
)

/*
实现strings.ReplaceAll函数
s:="abcabd" 中所有"ab"替换为"1"，返回结果"cd"
*/

func substr2(dest, s1, s2 string) []string {
	if len(dest) == 0 {
		return []string{}
	}

	n1 := len(dest)
	n2 := len(s1)

	res := []string{}
	i := 0
	for i < n1 {
		if i+n2 > n1 {
			break
		}
		s := dest[i : i+n2]
		if s == s1 {
			i = i + n2
			for j := 0; j < n2; j++ {
				res = append(res, s2)
			}
		} else {
			res = append(res, s)
			i++
		}
	}

	fmt.Println(res)
	result := []string{}
	for i := 0; i < len(res); i++ {
		if res[i] != s2 {
			result = append(result, res[i])
		}
	}
	return result
	//return strings.ReplaceAll(dest, s1, s2)
}

func main() {
	fmt.Println(substr2("abcabd", "ab", "1"))
}
