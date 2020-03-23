package main

import "fmt"

/*
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
*/

func hammingWeight(num uint32) int {
	move := uint32(1)
	count := 0
	for i := 0; i < 32; i++ {
		if (num & move) != 0 {
			count++
		}

		move = move << 1
	}

	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		count++
		num = num & (num - 1)
	}

	return count
}

func main() {
	//2^31 -1
	a := (1 << 31) - 1
	x := uint32(a)
	fmt.Println(hammingWeight2(x))
}
