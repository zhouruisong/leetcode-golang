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

//时间复杂度：O(S)，其中 SS 为所有可行解的长度之和， 上届是o(n*2^n)
// https://leetcode-cn.com/problems/combination-sum-ii/solution/man-tan-wo-li-jie-de-hui-su-chang-wen-shou-hua-tu-/
// 排序
func combinationSum2(candidates []int, target int) [][]int {
	/*
		通39题比较只需改动三点：
		给定的数组可能有重复的元素，先排序，使得重复的数字相邻，方便去重。
		for 枚举出选项时，加入下面判断，从而忽略掉同一层重复的选项，避免产生重复的组合。比如[1,2,2,2,5]，选了第一个 2，变成 [1,2]，
		它的下一选项也是 2，跳过它，因为如果选它，就还是 [1,2]。

		if (i - 1 >= start && candidates[i - 1] == candidates[i]) {
			continue;
		}
		当前选择的数字不能和下一个选择的数字重复，给子递归传i+1，避免与当前选的i重复。
		dfs(i + 1, temp, sum + candidates[i]);
	*/
	sort.Ints(candidates)
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

		//
		for i := start; i < len(candidates); i++ {
			//去重复
			if i-1 >= start && candidates[i-1] == candidates[i] {
				continue
			}

			if target < candidates[i] { //剪枝
				return
			}

			temp = append(temp, candidates[i])
			//当前选择的数字不能和下一个选择的数字重复，给子递归传i+1，避免与当前选的i重复
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

func main() {
	x := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(x, 8))
}
