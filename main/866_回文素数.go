package main

import (
	"fmt"
	"math"
	"time"
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

// 判断整数 n 是否是素数
func isSushu(n int) bool {
	if n <= 2 || n&1 == 0 {
		return n == 2
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
我们继续分析，其实质数还有一个特点，就是它总是等于 6x-1 或者 6x+1，其中 x 是大于等于1的自然数。
如何论证这个结论呢，其实不难。首先 6x 肯定不是质数，因为它能被 6 整除；其次 6x+2 肯定也不是质数，
因为它还能被2整除；依次类推，6x+3 肯定能被 3 整除；6x+4 肯定能被 2 整除。那么，就只有 6x+1 和 6x+5 (即等同于6x-1)
可能是质数了。所以循环的步长可以设为 6，然后每次只判断 6 两侧的数即可。
*/

func isPrime2(n int) bool {
	if n <= 3 {
		return n > 1
	}

	if n%6 != 1 && n%6 != 5 {
		return false
	}
	s := math.Sqrt(float64(n))
	for i := 5; i <= int(s); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func primePalindrome(N int) int {
	if N >= 8 && N <= 11 {
		return 11
	}

	y := 0
	for {
		//判断回文数
		tmp := N
		for tmp > 0 {
			y = y*10 + tmp%10
			tmp = tmp / 10
		}

		//是回文数在判断是否是素数(质数)
		if y == N && isSushu(N) {
			return N
		}
		y = 0
		N += 1

		//不存在长度为8的回文素数
		if N >= 10000000 && N < 100000000 {
			N = 100000000
		}
	}

	return -1
}

//一亿以内所有素数
func countPrimes() map[uint64]bool {
	var source []uint64 = []uint64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}

	N := len(source)
	j := 0
	for n := uint64(101); n < 10000; n += 2 {
		for j = 1; j < N; j++ {
			if n%source[j] == 0 {
				break
			}
		}

		if j == N {
			source = append(source, n)
		}
	}

	//fmt.Println(source)
	N = len(source)

	for n := uint64(10001); n < 100000000; n += 2 {
		for j = 1; j < N; j++ {
			if n%source[j] == 0 {
				break
			}
		}

		if j == N {
			source = append(source, n)
		}
	}

	fmt.Println(len(source))

	mp := make(map[uint64]bool)
	//for k := 0; k < len(source); k++ {
	//	mp[source[k]] = true
	//}
	return mp
}

func main() {
	start := time.Now()
	//fmt.Println(primePalindrome(9989900)) //100030001
	fmt.Println(primePalindrome(61023998))
	end := time.Since(start)
	fmt.Printf("countPrimes time cost = %v\n", end)
}
