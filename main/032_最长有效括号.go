package main

import (
	"container/list"
	"fmt"
	"sort"
)

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	stack := list.New()
	res := []int{}

	//  )(()())
	for k, v := range s {
		//为 '('的字符,将位置入栈,后面匹配使用
		if v == '(' {
			stack.PushFront(k)
			continue
		}

		//匹配stack的栈顶为 '(',存储位置和当前')'的位置
		if v == ')' && stack.Len() > 0 {
			res = append(res, stack.Remove(stack.Front()).(int))
			res = append(res, k)
		}
	}

	//fmt.Println(res)
	sort.Ints(res)

	//计算最长连续序列长度
	i := 0
	ans := 0
	ln := len(res)

	for i < ln {
		j := i
		for j < ln-1 && res[j+1] == res[j]+1 {
			j++
			ans = max(ans, j-i+1)
		}

		i = j + 1
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	s := ")(()())"
	fmt.Print(longestValidParentheses(s))
}
