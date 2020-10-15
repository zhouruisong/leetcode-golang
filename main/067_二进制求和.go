package main

import "fmt"

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/

func addBinary(a string, b string) string {
	n1 := len(a)
	n2 := len(b)

	if n1 == 0 && n2 != 0 {
		return b
	}

	if n1 != 0 && n2 == 0 {
		return a
	}

	if n1 == 0 && n2 == 0 {
		return ""
	}

	s := ""
	i := n1 - 1
	j := n2 - 1
	//pre:=uint8(0)
	jinwei := uint8(0)
	sum := uint8(0)
	for i >= 0 && j >= 0 {
		//这里还有优化的地方，
		sum = jinwei + (a[i] - '0') + (b[j] - '0')
		//cur := (a[i] - '0') ^ (b[j] - '0') ^ pre
		jinwei = sum / 2
		s = fmt.Sprintf("%v", sum%2) + s
		i--
		j--
	}

	if j >= 0 {
		for j >= 0 {
			sum = jinwei + (b[j] - '0')
			jinwei = sum / 2
			s = fmt.Sprintf("%v", sum%2) + s
			j--
		}
	}

	if i >= 0 {
		for i >= 0 {
			sum = jinwei + (a[i] - '0')
			jinwei = sum / 2
			s = fmt.Sprintf("%v", sum%2) + s
			i--
		}
	}

	if jinwei > 0 {
		s = fmt.Sprintf("%v", jinwei) + s
	}

	return s
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
