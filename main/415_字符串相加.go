package main

import "fmt"

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
注意：
num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/

func addStrings(num1 string, num2 string) string {
	i, j, carry, temp := len(num1)-1, len(num2)-1, uint8(0), uint8(0)
	n1, n2 := uint8(0), uint8(0)

	res := ""
	for i >= 0 || j >= 0 {
		if i >= 0 {
			n1 = num1[i] - '0'
		} else {
			n1 = 0
		}

		if j >= 0 {
			n2 = num2[j] - '0'
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

	return res
}

func reverse(nums []byte) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func addStrings2(num1 string, num2 string) string {
	var maxSlice, minSlice []byte
	if len(num1) > len(num2) {
		maxSlice = []byte(num1)
		minSlice = []byte(num2)
	} else {
		maxSlice = []byte(num2)
		minSlice = []byte(num1)
	}
	reverse(maxSlice)
	reverse(minSlice)
	var tag byte = 0
	var temp, tempMin byte
	for i := 0; i < len(maxSlice); i++ {
		if i >= len(minSlice) {
			tempMin = '0'
		} else {
			tempMin = minSlice[i]
		}
		temp = maxSlice[i] + tempMin - '0' + tag
		if temp > '9' {
			maxSlice[i] = temp - 10
			tag = 1
		} else {
			maxSlice[i] = temp
			tag = 0
		}
	}
	if tag == 1 {
		maxSlice = append(maxSlice, '1')
	}

	reverse(maxSlice)
	return string(maxSlice)
}

func main() {
	arr1 := "123445454545"
	arr2 := "45673232332232232"
	fmt.Println(addStrings(arr1, arr2))
}
