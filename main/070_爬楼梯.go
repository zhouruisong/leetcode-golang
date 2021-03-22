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
func climbStairs(N int) int {
	/*
		斐波那契数列
		自顶向下 递归思路f(n)：爬上n阶楼梯方法数 可以得出f(n) = f(n -1) + f(n-2)
	*/
	if N <= 1 {
		return N
	}

	one := 0
	two := 1
	result := 0
	for j := 2; j <= N; j++ {
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
	if n <= 1 {
		return n
	}

	return climbStairs2(n-1) + climbStairs2(n-2)
}

//动态规划
func climbStairs3(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//如果我们把问题泛化，不再是固定的1，2，而是任意给定台阶数，例如1,2,5呢？
//我们只需要修改我们的DP方程DP[i] = DP[i-1] + DP[i-2] + DP[i-5], 也就是DP[i] = DP[i] + DP[i-j] ,j =1,2,5
//在原来的基础上，我们的代码可以做这样子修改
func climbStairs4(n int, steps []int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	//steps := []int{1, 2, 5} // 不再是固定的1，2，而是任意给定台阶数，例如1,2,5
	for i := 1; i <= n; i++ {
		for j := 0; j < len(steps); j++ {
			step := steps[j]
			if i < step {
				continue
			}
			dp[i] = dp[i] + dp[i-step]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs4(3, []int{1, 2}))
}
