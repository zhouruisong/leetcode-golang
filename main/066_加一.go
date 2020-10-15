package main

import "fmt"

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {
	sum := 0
	n := len(digits)
	jinwei := 0
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			s := digits[i] + 1
			digits[i] = s % 10
			jinwei = s / 10
		} else {
			sum = jinwei + digits[i]
			jinwei = sum / 10
			digits[i] = sum % 10
		}
	}

	if jinwei > 0 {
		var ret []int
		ret = append(ret, jinwei)
		ret = append(ret, digits...)
		return ret
	} else {
		return digits
	}
}

func main() {
	digits := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
