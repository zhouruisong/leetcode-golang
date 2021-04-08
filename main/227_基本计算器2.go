package main

import "fmt"

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
示例 1：
输入：s = "3+2*2"
输出：7
示例 2：

输入：s = " 3/2 "
输出：1
示例 3：

输入：s = " 3+5 / 2 "
输出：5

提示：
1 <= s.length <= 3 * 105
s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
s 表示一个 有效表达式
表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
题目数据保证答案是一个 32-bit 整数
*/

func calculate(s string) int {
	//采用栈存放每个符号和值,遇到*和/时,当前值和栈顶元素作运算将结果入栈,遇到-号,当前值变为负数入栈
	stack := []int{}
	//sign := "+"
	num := 0
	for _, ch := range s {
		if ch == ' ' {
			continue
		}

		//"123+5"这样的数变为123 + 5 = 128
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}

		switch ch {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			n := stack[len(stack)-1] * num
			stack[len(stack)-1] = n
		case '/':
			n := stack[len(stack)-1] / num
			stack[len(stack)-1] = n
		default:
			stack = append(stack, num)
		}
		num = 0
	}

	result := 0
	for _, v := range stack {
		result += v
	}

	return result
}

func main() {
	s := "3+2*2"
	fmt.Println(calculate(s))
}
