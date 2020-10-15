package main

import (
	"fmt"
	"sort"
)

/* 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*
*
*
* 示例 1:
*
* 输入: [2,3,-2,4]
* 输出: 6
* 解释: 子数组 [2,3] 有最大乘积 6。
*
*
* 示例 2:
*
* 输入: [-2,0,-1]
* 输出: 0
* 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*
 */

func maxProduct(nums []int) int {
	/*
	 dp的思路就是把之前每个乘积运算记录下来，但是因为数组中存在负数 所有可能存在上一项是正数乘以当前负数成了更小的 ，
	 所以还需要记录一下上一项的最小值。也就是存在最大和最小值互相转换的可能性。
	 因此maxDP[i],minDP[i]= numCompare(minDP[i-1]*nums[i],maxDP[i-1]*nums[i],nums[i]) ,通过三者排序获得当前最大值和最小值。
	*/
	maxDP := make([]int, len(nums))
	minDP := make([]int, len(nums))
	maxDP[0] = nums[0]
	minDP[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i], minDP[i] = numCompare(minDP[i-1]*nums[i], maxDP[i-1]*nums[i], nums[i])
		if max < maxDP[i] {
			max = maxDP[i]
		}
	}
	return max
}

func numCompare(a, b, c int) (max, min int) {
	s := []int{a, b, c}
	sort.Ints(s)
	return s[2], s[0]
}

//第二种方法
func maxProduct2(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//arr := []int{2, 3, -2, 4}
	arr := []int{2, -5, -2, -4, 3}
	fmt.Println(maxProduct2(arr))
}
