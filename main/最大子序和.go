package main

import "fmt"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
func maxSubArray(x []int) ([]int, int) {
	start := 0 //开始索引
	end := 0   //结束索引
	s := 0
	e := 0
	//序列和
	sum := 0
	//最大和
	max := x[0]
	var ret []int
	ln := len(x)
	for i := 0; i < ln; i++ {
		sum += x[i]
		if sum > max {
			max = sum
			e = i
			start = s
			end = e
		}

		if sum < 0 {
			sum = 0
			s = i + 1
			e = i + 1
		}
	}

	//fmt.Printf("start=%+v，end=%+v\n", start, end)
	for i := start; i <= end; i++ {
		ret = append(ret, x[i])
	}
	return ret, max
}

func maxSubArray2(nums []int) int {
	sum := nums[0]
	max := 0
	start := 0
	end := 0

	n := len(nums)
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		if sum > max {
			max = sum
			end = i
		}

		if sum < 0 {
			sum = 0
			start = i + 1
			end = i + 1
		}
	}

	fmt.Printf("max:%v, start:%v, end:%v\n", max, start, end)
	return max
}

func main() {
	x := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//y, ret := maxSubArray(x)
	//fmt.Printf("y=%+v,ret=%+v\n", y, ret)

	maxSubArray2(x)
}
