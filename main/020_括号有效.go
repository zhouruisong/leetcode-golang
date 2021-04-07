package main

import (
	"fmt"
)

/*
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 */

func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	//栈
	qu := []string{}
	str := ""
	for i := 0; i < n; i++ {
		str = string(s[i])
		//fmt.Println("str: ", str)

		if str == "(" || str == "[" || str == "{" {
			qu = append(qu, str)
			//fmt.Println("push str: ", str)
			continue
		}

		if len(qu) > 0 {
			top := qu[len(qu)-1]
			//fmt.Println("top: ", top)

			if str == ")" && top == "(" {
				//fmt.Println("Remove str: ", str)
				qu = qu[:len(qu)-1]
				continue
			} else if str == "]" && top == "[" {
				//fmt.Println("Remove str: ", str)
				qu = qu[:len(qu)-1]
				continue
			} else if str == "}" && top == "{" {
				//fmt.Println("Remove str: ", str)
				qu = qu[:len(qu)-1]
				continue
			}
		}

		return false
	}

	if len(qu) > 0 {
		return false
	}
	return true
}

func main() {
	x := "([)]"
	//x := "()[]{}"
	//x := "]"
	//x := "["
	fmt.Println(isValid(x))
}
