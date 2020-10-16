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

func backtrace(start int, nums []int, res *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	//去重复
	m := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[start], nums[i] = nums[i], nums[start]
		backtrace(start+1, nums, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrace(0, nums, &res)
	return res
}

func permuteUnique2(nums []int) [][]int {
	//先将nums升序排序
	sort.Ints(nums)

	ln := len(nums)
	if ln == 0 {
		return [][]int{}
	}

	if ln == 1 {
		return [][]int{{nums[0]}}
	}

	res := [][]int{}
	//深拷贝
	tmp1 := make([]int, ln)
	copy(tmp1, nums)
	res = append(res, tmp1)

	//去重复
	m := make(map[string][]int)
	s := fmt.Sprintf("%v", tmp1)
	ret := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	m[ret] = tmp1

	for {
		//fmt.Println(nums)
		j := ln - 1
		k := ln - 1
		//从后往前找
		for j > 0 && nums[j-1] >= nums[j] {
			j-- //找到第一对数 nums[j]<nums[j+1]
		}

		if j < 0 {
			break
		}

		//去重复
		if j > 0 {
			j--
		}

		//从后往前找，到j位置结束，找到第一个比nums[j]大的nums[k]
		for k >= j && nums[k] <= nums[j] {
			k--
		}

		//去重复
		if k < 0 {
			break
		}

		//交换nums[j]和nums[k]对应的值
		nums[j], nums[k] = nums[k], nums[j]

		//j位置后，肯定是降序的，前后互换，变为升序
		l := j + 1
		r := ln - 1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			r--
			l++
		}

		//fmt.Printf("nums= %v\n", nums)

		s := fmt.Sprintf("%v", nums)
		ret = strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
		//fmt.Printf("%v\n", ret)

		//res中是否已经有了对于的[]int
		if _, ok := m[ret]; !ok {
			//深拷贝
			tmp := make([]int, ln)
			copy(tmp, nums)
			res = append(res, tmp)

			m[ret] = tmp
		}
	}

	return res
}

func main() {
	x := []int{1, 1, 5}
	fmt.Println(permuteUnique(x))
	y := []int{1, 1, 5}
	fmt.Println(permuteUnique(y))
}
