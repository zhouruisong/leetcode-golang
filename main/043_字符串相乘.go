package main

import "fmt"

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	if len(num1) == 1 && num1[0] == '0' {
		return "0"
	}

	if len(num2) == 1 && num2[0] == '0' {
		return "0"
	}

	n1 := len(num1)
	n2 := len(num2)
	N := n1
	arr := num1
	arr2 := num2
	if n2 <= n1 {
		N = n2
		arr = num2
		arr2 = num1
	}

	res := make([][]uint8, N)
	k := 0
	for i := N - 1; i >= 0; i-- {
		result1 := multiplyHelper(arr2, arr[i])
		for m := 0; m < k; m++ {
			result1 = append(result1, 0)
		}
		k++
		res[i] = result1
	}

	//每一行累加
	for i := 0; i < N-1; i++ {
		fmt.Println(res[i], res[i+1])
		res[i+1] = addStrings(res[i], res[i+1])
	}

	r := []byte{}
	for _, val := range res[N-1] {
		r = append(r, val+'0')
	}

	return string(r)
}

func addStrings(num1, num2 []uint8) []uint8 {
	fmt.Println(num1, num2)
	i, j, carry, temp := len(num1)-1, len(num2)-1, uint8(0), uint8(0)
	n1, n2 := uint8(0), uint8(0)

	res := ""
	for i >= 0 || j >= 0 {
		if i >= 0 {
			n1 = num1[i]
		} else {
			n1 = 0
		}

		if j >= 0 {
			n2 = num2[j]
		} else {
			n2 = 0
		}

		temp = n1 + n2 + carry
		carry = temp / 10
		res = fmt.Sprintf("%d%s", temp%10, res)
		i--
		j--
	}

	if carry == 1 {
		res = fmt.Sprintf("%d%s", carry, res)
	}

	byteArr := []byte(res)
	result := []uint8{}
	for i := 0; i < len(byteArr); i++ {
		result = append(result, byteArr[i]-'0')
	}
	return result
}

func multiplyHelper(arr string, x byte) []uint8 {
	n := len(arr)
	dx := x - '0'
	if dx == 0 {
		return []uint8{}
	}

	res := []uint8{}
	var jinwei uint8
	var yu uint8
	var v uint8

	for i := n - 1; i >= 0; i-- {
		v = dx*(arr[i]-'0') + jinwei
		yu = v % 10
		jinwei = v / 10
		res = append(res, yu)
	}

	if jinwei >= 1 {
		res = append(res, jinwei)
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return res
}

func main() {
	//num1 := "123"
	//num2 := "11"

	fmt.Println(multiply("733064366", "459309139"))
}
