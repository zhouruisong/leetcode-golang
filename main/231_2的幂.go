package main

import "fmt"

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 20 = 1
示例 2:

输入: 16
输出: true
解释: 24 = 16
示例 3:

输入: 218
输出: false
*/

//二的幂必须是 1<<n这个二进制数字
func isPowerOfTwo2(n uint64) bool {
	for n >= 2 {
		if n&1 == 1 {
			return false
		}
		n = n >> 1
	}

	return n == 1
}

/*
逆向操作，判断是不是2的幂，就看跟2的幂是否相等
比较n和2的幂的大小，若n < 1 (30) 则返回false，n ==1 返回true
n > 2 ,进入下次循环，比较 n < 2 (31)，依次循环
*/
func isPowerOfTwo3(n uint64) bool {
	aa := uint64(1)
	for {
		if n < aa {
			return false
		}

		if n == aa {
			return true
		}
		aa = aa * 2
	}
}

func main() {
	//fmt.Println(isPowerOfTwo2(256))

	res := []uint64{}
	for i := uint64(1); i < 1e9; i++ {
		if !isPowerOfTwo2(i) {
			continue
		}
		res = append(res, i)
	}

	fmt.Println(res)
}
