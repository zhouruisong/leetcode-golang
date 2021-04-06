package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

//o(n *n!)
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make(map[int]bool, len(nums))
	sort.Ints(nums) //要排序
	dfs22(nums, path, &res, used)
	return res
}

func dfs22(nums, path []int, res *[][]int, used map[int]bool) {
	if len(path) == len(nums) {
		//temp := make([]int, len(nums))
		//copy(temp, path)
		//*res = append(*res, temp)
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		//去重复
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		if !used[i] {
			//未使用过才进行
			used[i] = true
			path = append(path, nums[i])
			dfs22(nums, path, res, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func permuteUnique3(nums []int) [][]int {
	//字典序+map去重复
	sort.Ints(nums)

	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{nums[0]}}
	}
	res := [][]int{}
	tmp1 := make([]int, len(nums))
	copy(tmp1, nums)
	res = append(res, tmp1)
	//res = append(res, append([]int{}, nums...))

	//map去重复
	mp := make(map[string]struct{})
	s := fmt.Sprintf("%v", tmp1)
	key := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	mp[key] = struct{}{}

	for {
		//从后往前找，第一个连续的nums[j]<nums[j+1]
		j := n - 2
		for j >= 0 {
			if nums[j] < nums[j+1] {
				break
			}
			j--
		}
		if j < 0 {
			break
		}
		//从后到j找，满足第一个nums[k]>nums[j]
		k := n - 1
		for k >= j {
			if nums[k] > nums[j] {
				//交换
				nums[j], nums[k] = nums[k], nums[j]
				break
			}
			k--
		}

		if k < j {
			break
		}

		//j+1到最后变升序
		sort.Ints(nums[j+1:])
		//去重复
		s := fmt.Sprintf("%v", nums)
		key = strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
		if _, ok := mp[key]; !ok {
			//tmp := make([]int, len(nums))
			//copy(tmp, nums)
			//res = append(res, tmp)
			res = append(res, append([]int{}, nums...))
			mp[key] = struct{}{}
		}
	}
	return res
}

func main() {
	//x := []int{1, 1, 2}
	x := []int{2, 2, 1, 1}
	fmt.Println(permuteUnique(x))
	//y := []int{1, 1, 5}
	fmt.Println(permuteUnique3(x))
}
