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
*/

//栈
func largestRectangleArea(heights []int) int {
	s := list.New()
	heightsArr := []int{0}
	heightsArr = append(append(heightsArr, heights...), 0)
	//fmt.Println(heightsArr)

	res := float64(0)
	for k, v := range heightsArr {
		if k == 0 {
			s.PushFront(0)
		}

		curLen := s.Len()
		for curLen > 0 && heightsArr[s.Front().Value.(int)] > v {
			cur := s.Remove(s.Front()).(int)
			top := s.Front().Value.(int)

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
