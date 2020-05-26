package main

import (
	"fmt"
	"math"
)

/* 给定一个非负整数数组，你最初位于数组的第一个位置。
*
* 数组中的每个元素代表你在该位置可以跳跃的最大长度。
*
* 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*
* 示例:
*
* 输入: [2,3,1,1,4]
* 输出: 2
* 解释: 跳到最后一个位置的最小跳跃数是 2。
* 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*
*
* 说明:
*
* 假设你总是可以到达数组的最后一个位置。
 */

//O(n) 贪心算法 最优
func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		//找能跳的最远的（使用最少的跳跃次数）
		maxPosition = int(math.Max(float64(maxPosition), float64(nums[i]+i)))
		if i == end { //遇到边界，就更新边界，并且步数加一
			end = maxPosition
			steps++
		}
	}
	return steps
}

//动态规划
func jump(nums []int) int {
	// 动态规划四步走：
	// 1、状态：f[i] 表示从起点到当前位置跳的最小次数
	// 2、推导：f[i] = min(f[j]+1),a[j]+j >=i 表示从j位置用一步跳到当前位置，这个j位置可能有多个，取最小的一个就行
	// 3、初始化：f[0] = 0
	// 4、结果：f[n-1]
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		// f[i] 先取一个默认最大值i
		f[i] = i
		// 遍历之前结果取一个最小值+1
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = min(f[j]+1, f[i])
			}
		}
	}
	return f[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
