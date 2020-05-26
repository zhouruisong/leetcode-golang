package main

import (
	"fmt"
)

/*
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2:
 *
 * 输入: [2,7,9,3,1]
 * 输出: 12
 * 解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 */

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return Max(nums[0], nums[1])
	} else if n == 3 {
		return Max(nums[0]+nums[2], nums[1])
	} else {
		m := []int{}
		m = append(m, nums[0])
		m = append(m, Max(nums[0], nums[1]))
		m = append(m, Max(nums[0]+nums[2], nums[1]))
		for i := 3; i < n; i++ {
			//要想挖到第k家时获得金额的最大值，只需判断是从k-2家挖，还是k-3家挖即可
			m = append(m, Max(nums[i]+m[i-2], nums[i]+m[i-3]))
		}

		//fmt.Println(m)
		return Max(m[len(m)-1], m[len(m)-2])
	}
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	if nums[1] < nums[0] {
		nums[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		if nums[i-1] < nums[i-2]+nums[i] {
			nums[i] += nums[i-2]
		} else {
			nums[i] = nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

/*
动态规划方程：dp[n] = MAX( dp[n-1], dp[n-2] + num )
由于不可以在相邻的房屋闯入，所以在当前位置 n 房屋可盗窃的最大值，要么就是 n-1 房屋可盗窃的最大值，
要么就是 n-2 房屋可盗窃的最大值加上当前房屋的值，二者之间取最大值
举例来说：1 号房间可盗窃最大值为 3， 即为 dp[1]=3，2 号房间可盗窃最大值为 4， 即为 dp[2]=4，3 号房间自身的值为 2 即为 num=2，
那么 dp[3] = MAX( dp[2], dp[1] + num ) = MAX(4, 3+2) = 5，3 号房间可盗窃最大值为 55
时间复杂度：O(n)，n 为数组长度
*/

func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= n; i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//x := []int{2, 7, 9, 3, 1}
	x := []int{1, 2, 3, 1}

	fmt.Println(rob2(x))
}
