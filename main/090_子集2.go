package main

import (
	"fmt"
	"sort"
	"time"
)

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

//回溯法
func subsetshelper(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		//返回[[]]
		return append([][]int{}, []int{})
	}

	item := nums[n-1]
	res := subsetshelper(nums[:n-1])

	ln := len(res) //提前计算出ln，后面res的len会发生变化,防止出现死循环
	mp := make(map[string]bool)
	//为了去重复
	for i := 0; i < ln; i++ {
		mp[fmt.Sprintf("%v", res[i])] = true
	}

	for i := 0; i < ln; i++ {
		//此处要深拷贝
		tmp := make([]int, len(res[i]))
		copy(tmp, res[i])
		tmp = append(tmp, item)
		//去重复
		if mp[fmt.Sprintf("%v", tmp)] {
			continue
		}
		res = append(res, tmp)
		//mp[fmt.Sprintf("%v", tmp)] = true
	}

	return res
}

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	return subsetshelper(nums)
}

func main() {
	x := []int{1, 2, 2}
	start := time.Now()
	fmt.Println(subsetsWithDup(x))
	fmt.Print("cost time is: " + time.Since(start).String())
}
