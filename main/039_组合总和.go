package main

import (
	"fmt"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

//时间复杂度：O(S)，其中 S 为所有可行解的长度之和， 上届是o(n*2^n)
func combinationSum(candidates []int, target int) [][]int {
	// https://leetcode-cn.com/problems/combination-sum/solution/shou-hua-tu-jie-zu-he-zong-he-combination-sum-by-x/
	// 未排序

	res := [][]int{}
	var dfs func(start int, temp []int, sum int)

	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				//深拷贝
				//newTmp := make([]int, len(temp))
				//copy(newTmp, temp)
				//res = append(res, newTmp)
				res = append(res, append([]int{}, temp...))
			}
			return
		}

		//子递归传了 i 而不是 i+1 ，因为元素可以重复选入集合，如果传 i+1 就不重复了。
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

func main() {
	//x := []int{2, 3, 6, 7}
	//fmt.Println(combinationSum(x, 7))
	x := []int{2, 3, 5}
	fmt.Println(combinationSum(x, 8))
}
