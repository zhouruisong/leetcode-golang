package main

import (
	"fmt"
	"sort"
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

//o(n*2^n)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	} else if target < 0 {
		return
	} else {
		for i := left; i < len(candidates); i++ { // left限定，形成分支
			if target < candidates[i] { //剪枝
				return
			}

			dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res) //*分支
		}
	}
}

func main() {
	x := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(x, 8))
}
