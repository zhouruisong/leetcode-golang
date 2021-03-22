package main

import "fmt"

/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。



示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2:

输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。
示例 3:

输入: amount = 10, coins = [10]
输出: 1
*/

func change(amount int, coins []int) int {
	/*
		动态规划
		定义
		d[i] 表示总金额为i的时候按照目前coints可以构成的组合的数量
		这里故意用 d[amount+1] 因为要考虑d[0]的情况

		初始化
		d[0]=1 当金额为0，唯一组合就是0个金币，所以为1
		d[i]=0 i > 0

		计算

		外循环不同的金币 coin 来做遍历，
		内循环考虑金额 x=0~amount（这里可以从coin开始），加上之前x金额的数量 + x-coin金额的数量，就是总和，即 d[x] += d[x-coin]
		结果
		d[amount]

		额外的解释

		为什么只用一维d[i]
		按照标准背包做法，其实是用二维的，考虑它之和比它小编号的有关，所以可以用一维来解决
		另外一个潜在收益，通过 x>= coin来剪枝的遍历，减少了二维的赋值，速度进一步变快
		如何理解 d[x] += d[x-coin]
		二维 d[i][x] 表示对于i种硬币的情况下，总额为x的组合数量
		d[i+1][x] 表示多了 i+1种硬币的情况下，基于已经存在组合数量 d[i][x]， 加上使用了coin（第i+1枚硬币对应的价值）情况下
		d[i+1][x-coin] 的数量，总和就是 d[i][x] + d[i+1][x-coin]
		考虑到 0到n的情况，而且变为一维，那么就是 d[x] += d[x-coin] 对于每种硬币的情况
	*/
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for i := 0; i <= amount; i++ {
			if i >= c {
				dp[i] = dp[i] + dp[i-c]
			}
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
