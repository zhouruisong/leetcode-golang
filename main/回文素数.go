package main

import (
	"fmt"
	"math"
)

/*
求出大于或等于 N 的最小回文素数。

回顾一下，如果一个数大于 1，且其因数只有 1 和它自身，那么这个数是素数。

例如，2，3，5，7，11 以及 13 是素数。

回顾一下，如果一个数从左往右读与从右往左读是一样的，那么这个数是回文数。

例如，12321 是回文数。



示例 1：

输入：6
输出：7
示例 2：

输入：8
输出：11
示例 3：

输入：13
输出：101


提示：

1 <= N <= 10^8
答案肯定存在，且小于 2 * 10^8。
*/

//N是否是素数

func sushu(n int) bool {
	if n < 2 || n&1 == 0 {
		return n == 2
	}

	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primePalindrome(N int) int {
	y := uint64(0)
	limit := uint64(10 * 10 * 10 * 10 * 10 * 10 * 10 * 10)
	for i := uint64(1); i < limit; i++ {
		//判断回文数
		tmp := i
		for tmp != 0 {
			y = y*10 + tmp%10
			tmp = tmp / 10
		}

		//是回文数在判断是否是素数(质数)
		if i >=uint64(N) && y == i && sushu(int(i)) {
			fmt.Println(y)

			return int(i)
		}
		y = uint64(0)
	}

	return -1
}

func main() {
	//smp := make(map[uint64]uint64)
	//limit := uint64(10 * 10 * 10 * 10 * 10 * 10 * 10 * 10)
	//limit := uint64(10 * 10)
	//limit := uint64(20)

	//sushu(limit, smp)

	fmt.Println(primePalindrome(9989900))
	fmt.Println(primePalindrome(100))
}
