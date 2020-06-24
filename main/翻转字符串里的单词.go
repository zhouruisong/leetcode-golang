package main

import (
	"container/list"
	"fmt"
	"strings"
)

/* 给定一个字符串，逐个翻转字符串中的每个单词。
*
* 示例 1：
*
* 输入: "the sky is blue"
* 输出: "blue is sky the"
*
*
* 示例 2：
*
* 输入: "  hello world!  "
* 输出: "world! hello"
* 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
*
*
* 示例 3：
*
* 输入: "a good   example"
* 输出: "example good a"
* 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*
* 说明：
*
* 无空格字符构成一个单词。
* 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
* 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*
* 进阶：
*
* 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
*
 */

func reverseWords(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	arr := list.New()
	res := []byte{}
	i := 0
	for ; i < n; i++ {
		if s[i] == ' ' {
			if len(res) > 0 {
				arr.PushFront(string(res))
				//fmt.Println(string(res))
				res = []byte{}
			}
		} else {
			res = append(res, s[i])
		}
	}

	if i >= n && len(res) > 0 {
		arr.PushFront(string(res))
		//fmt.Println(string(res))
	}

	var ret []string
	for arr.Len() > 0 {
		node := arr.Remove(arr.Front()).(string)
		ret = append(ret, node)
	}

	return strings.Join(ret, " ")
}

func main() {
	s := "the sky is blue"
	//s := "  hello world!  "
	//s := "a good   example"
	//s := "  hello a"
	fmt.Println(reverseWords(s))
}
