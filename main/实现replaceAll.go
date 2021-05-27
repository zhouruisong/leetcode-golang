package main

import (
	"fmt"
	"strings"
	//"strings"
)

/*
实现strings.ReplaceAll函数
s:="abcabd" 中所有"ab"替换为"1"，返回结果"1c1d"
*/

func substr2(dest, s1, s2 string) string {
	if len(dest) == 0 {
		return ""
	}

	n1 := len(dest)
	n2 := len(s1)

	res := []string{}
	i := 0
	for i < n1 {
		if i+n2 > n1 {
			res = append(res, dest[i:])
			break
		}
		s := dest[i : i+n2]
		if s == s1 {
			i = i + n2
			res = append(res, s2)
		} else {
			res = append(res, string(dest[i]))
			i++
		}
	}

	result := ""
	for i := 0; i < len(res); i++ {
		result = result + res[i]
	}
	return result
}

func main() {
	fmt.Println(strings.ReplaceAll("abcabd", "ab", "1"))
	fmt.Println(substr2("abcabd", "ab", "1"))
}
