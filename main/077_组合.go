package main

import "fmt"

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

func combine2(n int, k int) [][]int {
	if k > n {
		return nil
	}
	var res [][]int
	var arr []int
	if n == 0 {
		return [][]int{{}}
	}
	getCombine2(1, k, &arr, &res, n)
	return res
}

func getCombine2(n int, k int, arr *[]int, res *[][]int, max int) {
	if len(*arr) == k {
		temp := make([]int, len(*arr))
		copy(temp, *arr)
		*res = append(*res, temp)
		return
	}
	if k-len(*arr) > max-n+1 {
		return
	}
	//下面屏蔽的代码是剪枝方式，但完全可以用上面的剪枝方式取代，而且上面更优。
	//if n>max {
	//	return
	//}
	getCombine2(n+1, k, arr, res, max)
	*arr = append(*arr, n)
	getCombine2(n+1, k, arr, res, max)
	res2 := *arr
	res2 = res2[:len(res2)-1]
	*arr = res2
}

func main() {
	fmt.Println(combine2(4, 2))
}
