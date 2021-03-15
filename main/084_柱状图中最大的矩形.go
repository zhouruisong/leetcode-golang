package main

import (
	"container/list"
	"fmt"
	"math"
)

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
示例:
输入: [2,1,5,6,2,3]

首先，要想找到第 i 位置最大面积是什么？
是以i 为中心，向左找第一个小于 heights[i] 的位置 left_i；向右找第一个小于于 heights[i] 的位置 right_i，即最大面积为 heights[i] * (right_i - left_i -1)
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/dong-hua-yan-shi-dan-diao-zhan-84zhu-zhu-03w3/
*/

//栈 o(2n)
func largestRectangleArea(heights []int) int {
	s := list.New()
	//前后都多加一个0
	heightsArr := []int{0}
	heightsArr = append(append(heightsArr, heights...), 0)
	//fmt.Println(heightsArr)

	//栈压入哨兵值，便于heights打头的数组进行操作
	//压入0为方便计算打头位置的面积

	res := float64(0)
	for k, v := range heightsArr {
		if k == 0 {
			s.PushFront(0)
		}

		curLen := s.Len()
		//栈里面后面比前面大的时候才压入，相当于顺序压入，当
		//当前值比栈顶的值小的时候，相当于两个比栈顶小的值把
		//栈顶位置的数卡在中间，比如5，6，2，栈顶数为6
		//此时可以计算栈顶6围成的矩形面积
		for curLen > 0 && heightsArr[s.Front().Value.(int)] > v {
			cur := s.Remove(s.Front()).(int)
			top := s.Front().Value.(int)
			//面积计算公式为当前下标值*(左右两边的坐标减去1)
			area := (k - top - 1) * heightsArr[cur]
			res = math.Max(float64(area), res)
			curLen--
		}

		s.PushFront(k)
	}

	return int(res)
}

func main() {
	arr := []int{2, 1, 5, 6, 2, 3}
	//arr := []int{2, 1, 2}

	fmt.Println(largestRectangleArea(arr))
}
