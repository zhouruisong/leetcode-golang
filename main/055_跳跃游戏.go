package main

/*
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 */

/*
第i个节点是否能够到达，取决于从0到i-1个节点是否有一个能跳到i节点
所以，dp[i] = {dp[0] || .... || dp[i-1]}
*/
//O(n)
func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < n; i++ {
		flag := false
		//数组中的每个元素代表你在该位置可以跳跃的最大长度。
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}

		dp[i] = flag
	}

	return dp[n-1]
}

/*
1.定义当前最远可达节点index
迭代中判断
2.最终点是否在当前点之前 是则中断出false
3.当前节点的最远可达节点（index+num）大于 历史最远可达节点 更新最远可达节点index
4.最远可达节点index 大于等于最后一位 返回true

时间复杂度 O(n)
空间复杂度 O(1)
*/
func canJump(nums []int) bool {
	length := len(nums)
	if nums == nil || length == 0 {
		return false
	}
	end := 0
	for i, num := range nums {
		if end < i {
			return false
		}
		thisMaxEnd := i + num
		if thisMaxEnd > end {
			end = thisMaxEnd
		}
		if end >= length-1 {
			return true
		}
	}
	return false
}
