package main

import "fmt"

/*
链接：https://www.nowcoder.com/questionTerminal/25b04bac36b840ab93f6fd913d1c7247
来源：牛客网

编程题
【问题描述】
给你一根长度为n的绳子，请把绳子剪成m段 (m和n都是整数，n>1并且m>1)每段绳子的长度记为k[0],k[1],...,k[m].请问k[0]*k[1]*...*k[m]可能的最大乘积是多少？
【要求】
【数据输入】绳子长度，一个整型数n(n>1)
【数据输出】把绳子剪成m段后, 可能的最大乘积数，也是一个整型数
【样例输入】
绳子的长度n=9
【样例输出】
我们把它剪成长度分别为3,3,3的三段，此时得到的最大乘积是3*3*3=27
*/

func cutRope(number int) int {
	if number <= 1 {
		return 1
	}
	//用一个dp数组记录0-number长度的绳子剪掉后的最大乘积
	dp := make([]int, number+1)
	//长度为2的绳子被剪只有1和1的两段,乘积结果为3
	dp[2] = 1

	maxf := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	//长度大于等于3的绳子才进入我们的讨论范围
	for i := 3; i < number+1; i++ {
		//对于长度为i的绳子,它可分为两段j,i-j
		//j的范围从2开始到i
		//不剪总长度乘积为j*(i-j)
		//剪的话长度乘积为j*dp[i-j]
		//取两者的最大值max(j*(i-j), j*dp[i-j])
		for j := 2; j < i; j++ {
			dp[i] = maxf(dp[i], maxf(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[number]
}

func main() {
	fmt.Println(cutRope(9))
}
