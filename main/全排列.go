package main

import (
	"fmt"
)

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

这段话是网上抄来：
1、首先看最后两个数4, 5。 它们的全排列为4 5和5 4, 即以4开头的5的全排列和以5开头的4的全排列。
由于一个数的全排列就是其本身，从而得到以上结果。
2、再看后三个数3, 4, 5。它们的全排列为3 4 5、3 5 4、 4 3 5、 4 5 3、 5 3 4、 5 4 3 六组数。
即以3开头的和4,5的全排列的组合、以4开头的和3,5的全排列的组合和以5开头的和3,4的全排列的组合.
从而可以推断，设一组数p = {r1, r2, r3, ... ,rn}, 全排列为perm(p)，pn = p - {rn}。
因此perm(p) = r1perm(p1), r2perm(p2), r3perm(p3), ... , rnperm(pn)。当n = 1时perm(p} = r1。
为了更容易理解，将整组数中的所有的数分别与第一个数交换，这样就总是在处理后n-1个数的全排列

举2个例子来说明
123的全排列
--首先遍历元素,然后把遍历到的每一个元素都和第一个元素交换
第一个和第一个交换 就是1和1交换 最后还是123
那么就形成了 123  213  321 这样的3组
(交换后 再换回来 还原成123 因为后面的交换都是在123的基础上交换的 所以swap要写2次)

--检查每一组除了第一个之外的剩余元素， 如果这些元素个数是2，
那么就对这剩下的2个元素全排列 就是123 132 ,  213 231 , 321 312
2个元素的全排列很简单 就是把这2个元素交换位置就OK）

1234的全排列
--首先遍历元素,然后把遍历到的每一个元素都和第一个元素交换
那么就形成了 1234  2134  3214 4231 这样的4组

--检查每一组除了第一个之外的剩余元素， 如1234剩余的是234，发现是3个元素
那么问题就转换为求234的全排列了
接下来也是一样 问题转换为求134， 214， 231的全排列

像这样 总是对除了第一个之外的元素全排列， 每次元素的个数都在减少一个，
求N个元素全排列最终就变成求2的元素的全排列了
*/

func permute2(nums []int) [][]int {
	ln := len(nums)
	if ln == 0 {
		return [][]int{}
	}

	if ln == 1 {
		return [][]int{{nums[0]}}
	}

	if ln == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		numsSubOne := append(numsCopy[:i], numsCopy[i+1:]...)
		valueSlice := []int{nums[i]}
		newSubSlice := permute2(numsSubOne)
		for _, newValue := range newSubSlice {
			result = append(result, append(valueSlice, newValue...))
		}
	}

	return result
}

func backtrace(start int, nums []int, res *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrace(start+1, nums, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permute1(nums []int) [][]int {
	res := [][]int{}
	backtrace(0, nums, &res)
	return res
}

func main() {
	x := []int{1, 1, 5}
	fmt.Println(permute1(x))
}
