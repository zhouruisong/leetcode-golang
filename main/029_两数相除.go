package main

import (
	"fmt"
	"math"
)

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2
*/

/*
不能用乘法，咱们就用加法和移位，总之都能实现乘2的功能
思路如下：
1）首先注意越界问题，被除数范围如下
[−2^31,  2^31 − 1]
当被除数和除数都是负数的时候，被除数最小可以为−2^31，当除数为-1时，得到的结果为2^31就越界了

2）然后我们以dividend=23 divisor=3来举例

第一步:
3 * 1 = 3 < 23
3 * 2 = 6 < 23
3 * 2 * 2 = 12 < 23
3 * 2 * 2 * 2 = 24 > 23
也就是说，divendend只能容下的divisor数目为
4 < divendend / divisor < 8

第二步:
我们把刚才能容下的4个divisor（即12）减掉，看看剩下的11还能存下几个divisor
这里的做法和第一步类似
3 * 1 = 3 < 11
3 * 2 = 6 < 11
3 * 2 * 2 = 12 > dividend
也就是说，11只能容下的divisor数目为
2 < 11 / divisor < 4

第三步:
我们把刚才能容下的2个divisor（即3）减掉，看看剩下的5还能存下几个divisor
这里的做法和第一步类似
3 * 1 = 3 < 5
3 * 2 = 6 > 5
也就是说，5只能容下的divisor数目为
1 < 5 / divisor < 2

第四步:
我们把刚才能容下的1个divisor（即6）减掉，看看剩下的2还能存下几个divisor
这时候被除数2已经比divisor小了，故而容不下任何divisor了

所以咱们divendend总共能容纳的divisor为：
4 + 2 + 1= 7
*/

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}

	// 注意越界问题
	if divisor == -1 && dividend == -(math.MaxInt32+1) {
		return math.MaxInt32
	}

	//取数的绝对值
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))
	result := div(dividendAbs, divisorAbs)

	// 还原其本来正负
	if (dividend <= 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		return -result
	}
	return result
}

func div(dividend int, divisor int) int {
	result := 0
	for dividend >= divisor {
		multi := 1
		for multi*divisor <= (dividend >> 1) {
			multi <<= 1
		}
		result += multi
		// 相减的结果进入下次循环
		dividend -= multi * divisor
	}
	return result
}

func main() {
	//fmt.Print(divide(23, 3))
	fmt.Print(divide(10, 3))
}
