package main

import "fmt"

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

"""
时间复杂度为O(n)， 空间复杂度为O(1)
思路：left、right游标分别从列表左右两端向中间靠拢
1、计算以left、right为左右游标的容量（取游标指向的值中较小的作为容器高度）
2、比较left、right两个游标指向值的大小，较小的往下一个位置移动，
否则随便选择一个游标下移，在本程序中固定选择右边的游标下移
3、重复步骤1 的计算，直到程序结束
"""
*/

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		l := left
		r := right
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left += 1
		} else {
			h = height[right]
			right -= 1
		}

		tmp := h * (r - l)
		if tmp > max {
			max = tmp
		}
	}

	fmt.Printf("max:%+v\n", max)
	return max
}

func main() {
	x := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(x)
}
