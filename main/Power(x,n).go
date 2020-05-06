package main

import "fmt"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

/*
O(logN)
那么假设n=11，11的二进制是1011，11 = 2³×1 + 2²×0 + 2¹×1 + 2º×1=2³+2¹+2º，所以：a¹¹= a^2º* a ^2¹ * a^2³
代码中n&1是取末位，只有当前位为1时才要乘； n/=2是将n右移一位，取新的位做末位；x*=x就是X^(2^i)，是下一次要乘的因子
*/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	pow := float64(1)
	for n != 0 {
		if n&1 == 1 {
			pow *= x
		}
		x = x * x
		n >>= 1 //相当于n/2
	}

	return pow
}

func main() {
	var x float64
	x = 0.00001
	fmt.Println(myPow(x, 2))
}
