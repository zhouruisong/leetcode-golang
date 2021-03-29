package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

比39题增加的步骤，已用*号提示
1.排序
2.深度优先搜索
3.*避免重复
3.剪枝
4.结算
*/

//o(n*2^n)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		//tmp := make([]int, len(nums))
		//copy(tmp, nums)
		//*res = append(*res, tmp)
		*res = append(*res, append([]int{}, nums...))
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if i != left && candidates[i] == candidates[i-1] { // *同层节点 数值相同，剪枝
			continue
		}
		if target < candidates[i] { //剪枝
			return
		}

		//作选择
		nums = append(nums, candidates[i])
		dfs(candidates, nums, target-candidates[i], i+1, res) //*分支 i+1避免重复
		//撤销选择
		nums = nums[:len(nums)-1]
	}
}

func main() {
	x := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(x, 8))
}
