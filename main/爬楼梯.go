package main

import "fmt"

/*
*
* 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
*
* 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*
* 注意：给定 n 是一个正整数。
*
* 示例 1：
*
* 输入： 2
* 输出： 2
* 解释： 有两种方法可以爬到楼顶。
* 1.  1 阶 + 1 阶
* 2.  2 阶
*
* 示例 2：
*
* 输入： 3
* 输出： 3
* 解释： 有三种方法可以爬到楼顶。
* 1.  1 阶 + 1 阶 + 1 阶
* 2.  1 阶 + 2 阶
* 3.  2 阶 + 1 阶
*
 */

//非递归
func climbStairs(n int) int {
	/*
	斐波那契数列
	自顶向下 递归思路f(n)：爬上n阶楼梯方法数 可以得出f(n) = f(n -1) + f(n-2)
	*/
	if n <= 2 {
		return n
	}

	one := 1
	two := 2
	result := 0
	for j := 3; j <= n; j++ {
		result = one + two
		one = two
		two = result
	}
	return result
}

//递归方法(耗时长)
func climbStairs2(n int) int {
	/*
	斐波那契数列
	自顶向下 递归思路f(n)：爬上n阶楼梯方法数 可以得出f(n) = f(n -1) + f(n-2)
	*/
	if n <= 2 {
		return n
	}

	return climbStairs2(n-1) + climbStairs2(n-2)
}

func main() {
	fmt.Println(climbStairs(44))
	//fmt.Println(climbStairs2(44))
}
