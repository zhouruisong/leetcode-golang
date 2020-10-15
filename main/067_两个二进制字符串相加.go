package main

import (
	"container/list"
	"fmt"
)

/*题目：
input a = "100", b="010"
output "110"

Input: a = "11", b = "1"
Output: "100"

Input: a = "1010", b = "1011"
Output: "10101"

解题思路：

采用10进制的方法相加，从低位加，超过2 向高位进位
*/

func binarySum(a, b string) {
	var out list.List
	a_len := len(a)
	b_len := len(b)

	fmt.Printf("a=%+v\n", a)
	fmt.Printf("b=%+v\n", b)

	i := a_len - 1
	j := b_len - 1

	//记录两个数的和
	sum := uint8(0)
	for i >= 0 && j >= 0 {
		tmpa := a[i] - 48
		tmpb := b[j] - 48
		sum = sum + tmpa + tmpb //记得加上进位数
		yu := sum % 2           //本位的值
		//fmt.Printf("sum=%+v, yu=%+v\n", sum, yu)

		out.PushFront(fmt.Sprintf("%v", yu))
		sum = sum / 2 //进位数

		i--
		j--
	}

	if sum > 0 {
		out.PushFront(fmt.Sprintf("%v", sum))
	}

	item := out.Front()
	for item != nil {
		fmt.Println(item.Value)
		item = item.Next()
	}
	return
}

func main() {
	//a := "100"
	//b := "010"

	a := "1010"
	b := "1011"

	binarySum(a, b)
}
