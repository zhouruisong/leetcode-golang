package main

import (
	"fmt"
	"time"
)

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

/*
第一个解法是利用数学归纳的思想：假设我现在知道了规模更小的子问题的结果，如何推导出当前问题的结果呢？
具体来说就是，现在让你求 [1,2,3] 的子集，如果你知道了 [1,2] 的子集，
是否可以推导出 [1,2,3] 的子集呢？先把 [1,2] 的子集写出来瞅瞅：
[ [],[1],[2],[1,2] ]
你会发现这样一个规律：
subset([1,2,3]) - subset([1,2]) = [3],[1,3],[2,3],[1,2,3]
而这个结果，就是把 sebset([1,2]) 的结果中每个集合再添加上 3。
换句话说，如果 A = subset([1,2]) ，那么：
subset([1,2,3]) = A + [A[i].add(3) for i = 1..len(A)]
这就是一个典型的递归结构嘛，[1,2,3] 的子集可以由 [1,2] 追加得出，[1,2] 的子集可以由 [1] 追加得出，
base case 显然就是当输入集合为空集时，输出子集也就是一个空集。
*/

//回溯法O(n2^n)
func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		//返回[[]]
		return append([][]int{}, []int{})
	}

	item := nums[n-1]
	res := subsets(nums[:n-1])

	ln := len(res) //提前计算出ln，后面res的len会发生变化,防止出现死循环
	for i := 0; i < ln; i++ {
		//此处要深拷贝
		tmp := make([]int, len(res[i]))
		copy(tmp, res[i])
		tmp = append(tmp, item)
		res = append(res, tmp)
		//fmt.Println(res)
	}

	return res
}

func main() {
	//x := []int{1, 2, 2}
	start := time.Now()
	x := []int{9, 0, 3, 5, 7}
	subsets(x)

	fmt.Println("cost time is: " + time.Since(start).String())
	//y := [][]int{}
	//y = append(y, []int{})
	//fmt.Println(y)
}
