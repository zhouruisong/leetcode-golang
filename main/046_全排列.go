package main

import (
	"fmt"
	"reflect"
	"sort"
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

//o(n^2)
func permute(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make(map[int]struct{}, len(nums))
	dfs3(nums, path, &res, used)
	return res
}

//深度优先遍历,获取所有路径,包括重复的o(n!*n)
func dfs3(nums, path []int, res *[][]int, used map[int]struct{}) {
	if len(path) == len(nums) {
		//temp := make([]int, len(nums))
		//copy(temp, path)
		//*res = append(*res, temp)
		*res = append(*res, append([]int{}, path...))
		return
	}

	for _, v := range nums {
		//优化
		if _, ok := used[v]; ok {
			continue
		}

		used[v] = struct{}{}
		path = append(path, v)
		dfs3(nums, path, res, used)
		path = path[:len(path)-1]
		delete(used, v)
	}
}

func permute2(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			//temp := make([]int, len(path))
			//copy(temp, path)
			res = append(res, append([]int{}, path...))
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

//遍历array查询元素是否存在
func isExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == value {
				return true
			}
		}
	}
	return false
}

//每次都排序,二分查找,查询元素是否存在
func SortInIntSlice(haystack []int, needle int) bool {
	sort.Ints(haystack)

	index := sort.SearchInts(haystack, needle)
	return index < len(haystack) && haystack[index] == needle
}

//如果是每次数据相同,首次排序一次,二分查找,后续不需要在排序,查询元素是否存在
func InIntSliceSortedFunc(haystack []int) func(int) bool {
	sort.Ints(haystack)

	return func(needle int) bool {
		index := sort.SearchInts(haystack, needle)
		return index < len(haystack) && haystack[index] == needle
	}
}

//如果是每次数据相同,采用map,查找,查询元素是否存在
func InIntSliceMapKeyFunc(haystack []int) func(int) bool {
	set := make(map[int]struct{})

	for _, e := range haystack {
		set[e] = struct{}{}
	}

	return func(needle int) bool {
		_, ok := set[needle]
		return ok
	}
}

//字典序实现
func permute4(nums []int) [][]int {
	sort.Ints(nums)
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

	res := [][]int{}
	//深拷贝
	//tmp := make([]int, ln)
	//copy(tmp, nums)
	res = append(res, append([]int{}, nums...))

	//123 后一个排列是132
	for {
		//从后到前查找第一个符合nums[j] < nums[j+1]的位置j
		j := ln - 2
		for j >= 0 && nums[j] > nums[j+1] {
			j--
		}
		if j < 0 {
			break
		}

		//从后到j位置查找第一个符合nums[k] > nums[j]的位置k
		k := ln - 1
		for k > j && nums[k] < nums[j] {
			k--
		}

		if k == j {
			break
		}

		//交换nums[j]和nums[k]对应的值
		nums[j], nums[k] = nums[k], nums[j]

		//从j位置到结尾一定是降序的,变为升序
		sort.Ints(nums[j+1:])

		//深拷贝结果
		//tmp := make([]int, ln)
		//copy(tmp, nums)
		//res = append(res, tmp)
		res = append(res, append([]int{}, nums...))
	}

	return res
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(permute(x))
	fmt.Println(permute4(x))

	//fmt.Println(permute1(x))

	//相同的切片修改后，重复append会修改原来的切片
	//res := [][]int{}
	//arr := []int{1, 2, 3}
	//arr[0], arr[2] = arr[2], arr[0]
	//fmt.Println(arr) // [3 2 1]
	//
	//res = append(res, arr)
	//fmt.Println(res) // [3 2 1]
	//
	//arr[1], arr[2] = arr[2], arr[1]
	//fmt.Println(arr) // [3 1 2]
	//res = append(res, arr)
	//
	//fmt.Println(res) // [[3 1 2] [3 1 2]]   ?
}
