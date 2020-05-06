package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"

给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
*/

//方法1，列出所有全排列，在排序，取出前面第k-1个
func getPermutation(n int, k int) string {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	res := permute(nums, k)
	if len(res) < k {
		s := fmt.Sprint(res[len(res)-1])
		return strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	}

	//slice排序
	var m myvalue2
	m = res
	sort.Sort(m)

	//slice转字符串
	s := fmt.Sprint(m[k-1])
	return strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
}

func backtrace(start, k int, nums []int, res *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrace(start+1, k, nums, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permute(nums []int, k int) [][]int {
	res := [][]int{}
	backtrace(0, k, nums, &res)
	return res
}

type myvalue2 [][]int

func (p myvalue2) Len() int {
	return len(p)
}

// 想了想能用循环，实现所有字段比对一次，更精确点，免得二维数组的元素前面n个元素都是一样的
func (p myvalue2) Less(i, j int) bool {
	for k := 0; k < len(p[i]); k++ {
		if p[i][k] > p[j][k] {
			return false
		} else if p[i][k] == p[j][k] {
			continue
		} else {
			return true
		}
	}
	return true
}

func (p myvalue2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//方法很妙!
func getPermutation2(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n)
	//求0-9的阶乘
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i
		tokens[i-1] = i
	}

	k--
	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}

func getPermutation3(n int, k int) string {
	if n > 9 || n <= 0 {
		return ""
	}

	s := []byte("123456789")
	result := make([]byte, n)

	var factorial = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880} //阶乘0-10
	if k > factorial[n] {
		return ""
	}

	//关键步骤
	k = k - 1
	for i := n - 1; i >= 0; i-- {
		id := k / factorial[i]
		k = k - factorial[i]*id

		result[n-i-1] = s[id]
		s = append(s[:id], s[id+1:]...)
	}

	return string(result)
}

func main() {
	fmt.Println(getPermutation2(8, 20545))
}
