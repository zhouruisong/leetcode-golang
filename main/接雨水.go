package main

import (
	"fmt"
	"math"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/

func trap(height []int) int {
	//找出左边最大的和右边最大的
	ln := len(height)
	m_left := make([]float64, ln)
	m_right := make([]float64, ln)
	m_left[0] = float64(height[0])
	m_right[ln-1] = float64(height[ln-1])

	//遍历找出每个数的左边最大的值
	for i := 1; i < ln; i++ {
		m_left[i] = math.Max(m_left[i-1], float64(height[i]))
	}

	//遍历找出每个数的右边最大的值
	for i := ln - 2; i > 0; i-- {
		m_right[i] = math.Max(m_right[i+1], float64(height[i]))
	}

	fmt.Printf("m_left\n")

	for i := 0; i < ln; i++ {
		fmt.Printf("%+v,", m_left[i])
	}

	fmt.Printf("\nm_right\n")

	for i := 0; i < ln; i++ {
		fmt.Printf("%+v,", m_right[i])
	}

	//每个槽的水位=min(m_left[i], m_right[i]) - height[i]
	ret := 0
	for i := 0; i < ln; i++ {
		ret = ret + int(math.Min(m_left[i], m_right[i])-float64(height[i]))
	}

	fmt.Printf("\nret=:%+v\n", ret)
	return ret
}

//先找到从左到右第一个最大的高度的索引 a（可能有多个相同的最大高度），从左开始到这个位置[0,a]，下降位置才积水
//再从右向左找到第一个最大的高度的索引 b（（可能有多个相同的最大高度），从右开始到这个位置[b,n-1]，下降位置才积水
//才考虑两个位置的中间数据从开始到这个位置[a,b]，下降位置才积水
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	sum := 0 //接雨水量
	left_max_index := 0
	right_max_index := 0

	//从左侧向右侧找到第一个最大高度的位置
	tmp1 := height[0]
	for i := 1; i < n; i++ {
		//fmt.Println(height[i], tmp1)
		if height[i] > tmp1 {
			left_max_index = i
			tmp1 = height[i]
			//fmt.Println(left_max_index, tmp1)
		}
	}

	fmt.Printf("left_max_index: %v\n", left_max_index)

	tmp1 = height[n-1]
	//从右侧向左侧找到第一个最大高度的位置
	for k := n - 1; k >= 0; k-- {
		if height[k] >= tmp1 {
			right_max_index = k
			tmp1 = height[k]
			//fmt.Println(right_max_index, tmp1)
		}
	}

	fmt.Printf("right_max_index: %v\n", right_max_index)

	tmp1 = height[0]
	//从左侧开始向上爬，一直到left_max_index位置
	for j := 1; j <= left_max_index; j++ {
		//只有下降部分才能积水
		if height[j] < tmp1 {
			sum = sum + tmp1 - height[j]
		} else {
			tmp1 = height[j]
		}
	}

	tmp1 = height[n-1]
	//从右侧开始向上爬，一直到left_max_index位置
	for p := n - 2; p >= right_max_index; p-- {
		//只有下降部分才能积水
		if height[p] < tmp1 {
			sum = sum + tmp1 - height[p]
		}else{
			tmp1 = height[p]
		}
	}

	//查看中间部分从【left_max_index right_max_index]
	tmp1 = height[left_max_index]
	for q := left_max_index + 1; q <= left_max_index; q++ {
		if height[q] < tmp1 {
			sum = sum + tmp1 - height[q]
		} else {
			tmp1 = height[q]
		}
	}

	//fmt.Println(sum)
	return sum
}

func main() {
	//x := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//x := []int{4, 2, 0, 3, 2, 5}
	x := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	//trap(x)
	fmt.Println(trap2(x))
}
