package main

import (
	"fmt"
)

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	var result [][]int
	if k == 0 {
		return result
	}
	var nums = make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	helper(&result, nums, 0, k)
	return result
}

func helper(result *[][]int, nums []int, index, k int) {
	//达到目的，进行回溯
	if index == k {
		*result = append(*result, append([]int{}, nums[0:k]...))
		return
	}
	for start := index; start < len(nums); start++ {
		//可以在这里进行剪枝。去除递减的序列，防止结果重复
		if index == 0 || nums[start] > nums[index-1] {
			nums[start], nums[index] = nums[index], nums[start]
			helper(result, nums, index+1, k)
			nums[index], nums[start] = nums[start], nums[index]
		}
	}
	return

}

func combine3(n int, k int) [][]int {
	var result [][]int
	if k == 0 {
		return result
	}

	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	path := []int{}
	dfs(nums, path, &result, n, k, 0)
	return result
}

//N叉树的遍历框架
func dfs(arr, path []int, res *[][]int, n, k, start int) {
	//"终止条件" 找到一对组合
	if k == 0 {
		//拷贝结果
		//temp := make([]int, len(path))
		//copy(temp, path)
		//*res = append(*res, temp)
		*res = append(*res, append([]int{}, path...))
		return
	}

	//注意这里的i不能从0开始，如果从0开始会出现重复的，比如[1，2]和[2，1]
	for i := start; i < n-k+1; i++ {
		//逻辑运算1，（可有可无，视情况而定）
		//把当前值添加到集合中
		path = append(path, arr[i])
		//递归调用，遍历每一个分支
		dfs(arr, path, res, n, k-1, i+1)
		//逻辑运算2，（可有可无，视情况而定）
		//从当前分支跳到下一个分支的时候要把之前添加的值给移除
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine3(4, 2))
}
