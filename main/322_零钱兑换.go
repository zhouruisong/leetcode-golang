package main

import (
	"fmt"
	"math"
)

/* 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 *
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

func coinChange(coins []int, amount int) int {
	/*假设最小组合中最后一枚硬币是c：
	c 可能是 coins 中任何一个；
	去除 c 后剩下的部分，一定也是当前总额的一个最小组合（否则加上c不可能构成最小组合）
	或者用以下思路：
	如果 dp[i] 表示组成金额i的最小组合，dp[i]+1 一定是组成金额 i+c 的最小组合。
	假设我们知道F(S) ，即组成金额S最少的硬币数，最后一枚硬币的面值是C。那么由于问题的最优子结构，转移方程应为：
	F(S)=F(S−C)+1
	但我们不知道最后一枚硬币的面值是多少，所以我们需要枚举每个硬币面额值 c0, c1, c2 ...c{n -1}并选择其中最小值

	例子1：假设
	coins = [1, 2, 5], amount = 11
	则，当 i==0i==0 时无法用硬币组成，为 0 。当 i<0i<0 时，忽略 F(i)F(i)

	F(i)	最小硬币数量
	F(0)	0 //金额为0不能由硬币组成
	F(1)	1 //F(1)=min(F(1-1),F(1-2),F(1-5))+1=1F(1)=min(F(1−1),F(1−2),F(1−5))+1=1
	F(2)	1 //F(2)=min(F(2-1),F(2-2),F(2-5))+1=1F(2)=min(F(2−1),F(2−2),F(2−5))+1=1
	F(3)	2 //F(3)=min(F(3-1),F(3-2),F(3-5))+1=2F(3)=min(F(3−1),F(3−2),F(3−5))+1=2
	F(4)	2 //F(4)=min(F(4-1),F(4-2),F(4-5))+1=2F(4)=min(F(4−1),F(4−2),F(4−5))+1=2
	...	...
	F(11)	3 //F(11)=min(F(11-1),F(11-2),F(11-5))+1=3F(11)=min(F(11−1),F(11−2),F(11−5))+1=3
	我们可以看到问题的答案是通过子问题的最优解得到的。
	*/
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}

	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	//dp[i]表示凑到i元,需要的最少硬币个数, c是最后一个硬币
	dp := make([]int, amount+1)
	dp[0] = 0
	minF := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		//求范围内,dp[i]的最小值
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = minF(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func main() {
	arr := []int{1, 2, 5, 7, 10}
	fmt.Println(coinChange(arr, 14))
}
