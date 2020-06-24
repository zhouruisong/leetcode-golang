package main

import "fmt"

/*
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
此题我们要保证生成的括号字符串合法。
我们用l，r分别记录可以插入 '(' 和 ')' 的数量。
例如n = 3，那么一开始r = 0，l = 3。
当我们在递归函数中，选择插入 '(' 时，l要 - 1，r要 + 1。
因为你插入你一个'('势必要在接下来插入一个')'。
当我们在递归函数中，选择插入 ')' 时，r只需- 1即可
*/
func generateParenthesis(n int) []string {
	res := []string{}
	generateParenthesisHelper(&res, n, 0, "")
	return res
}

func generateParenthesisHelper(res *[]string, l, r int, str string) {
	if l == 0 && r == 0 {
		*res = append(*res, str)
	}

	if l > 0 {
		generateParenthesisHelper(res, l-1, r+1, str+"(")
	}

	if r > 0 {
		generateParenthesisHelper(res, l, r-1, str+")")
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
