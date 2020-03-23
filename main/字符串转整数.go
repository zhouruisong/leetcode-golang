package main

import "fmt"

/*
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−2^31,  2^31 − 1]。如果数值超过这个范围，请返回  INT_MAX
 * (2^31 − 1) 或 INT_MIN (−2^31) 。
 *
 * 示例 1:
 *
 * 输入: "42"
 * 输出: 42
 *
 *
 * 示例 2:
 *
 * 输入: "   -42"
 * 输出: -42
 * 解释: 第一个非空白字符为 '-', 它是一个负号。
 * 我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
 *
 *
 * 示例 3:
 *
 * 输入: "4193 with words"
 * 输出: 4193
 * 解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
 *
 *
 * 示例 4:
 *
 * 输入: "words and 987"
 * 输出: 0
 * 解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
 * ⁠    因此无法执行有效的转换。
 *
 * 示例 5:
 *
 * 输入: "-91283472332"
 * 输出: -2147483648
 * 解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
 * 因此返回 INT_MIN (−2^31) 。
 *
 */
func myAtoi(str string) int {
	n := len(str)
	flag := true
	if n == 0 {
		return 0
	}
	INN_Min := -(1 << 31)
	INN_Max := (1 << 31) - 1
	i := 0
	var ret int
	for ;i<n && str[i] == ' '; {
		i++
	}

	if i >=n {
		return 0
	}

	if str[i] == '-' {
		flag = false
		i++
	} else if str[i] == '+' {
		flag = true
		i++
	}

	if i >=n {
		return 0
	}

	for ; i < n; i++{
		if (str[i]-'0' >= 0) && (str[i]-'0' <= 9) {
			//fmt.Println(str[i])
			ret = ret*10 + int(str[i]-'0')
			if ret > INN_Max && !flag {
				return INN_Min
			} else if ret > INN_Max && flag {
				return INN_Max
			}
		} else {
			break
		}
	}

	if !flag {
		return -ret
	}
	return ret
}

func main() {
	//x := "-91283472332"
	//x := "+-2"
	//x := "4193 with words"
	x:=" "
	fmt.Println(myAtoi(x))
}
