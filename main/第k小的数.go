package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

//第k小的数
func findKthMin(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	//转换成求第len(nums)-k+1大的数
	return findKthLargest(nums, len(nums)-k+1)
}

//O(n)
func findKthLargest(nums []int, k int) int {
	//快排序思想
	return partition(nums, k)
}

func partition(nums []int, k int) int {
	i := 0
	j := len(nums) - 1
	loc := nums[i]
	for i < j {
		for i < j && nums[j] >= loc {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

		for i < j && nums[i] <= loc {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i] = loc

	if len(nums)-i == k {
		return nums[i]
	}
	if len(nums)-i > k {
		return partition(nums[i+1:], k)
	}
	return partition(nums[:i], k+i-len(nums))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组

题目描述
输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。
示例1
输入
复制
[4,5,1,6,2,7,3,8],4
返回值
复制
[1,2,3,4]

*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	//找出第1小的数，第2小的数...第k小的数即可
	if k <= 0 {
		return []int{}
	}

	res := []int{}
	for i := 1; i <= k; i++ {
		res = append(res, findKthLargest(input, len(input)-i+1))
	}
	return res
}

//o(nlogk)
func GetLeastNumbers_Solution2(input []int, k int) []int {
	// write code here
	//找出第1小的数，第2小的数...第k小的数即可
	if k <= 0 || k > len(input) {
		return []int{}
	}

	h := NodeHeap{}
	heap.Init(&h)
	for i := 0; i < k; i++ {
		heap.Push(&h, input[i])
	}

	for i := k; i < len(input); i++ {
		top := heap.Pop(&h).(int)
		if input[i] < top {
			heap.Push(&h, input[i])
		} else {
			heap.Push(&h, top)
		}
	}
	//倒序
	sort.Ints(h)
	return h
}

type NodeHeap []int

func (h NodeHeap) Len() int { return len(h) }

// 大根堆
func (h *NodeHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *NodeHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
给定String类型的数组strArr，再给定整数k，请严格按照排名顺序打印 出次数前k名的字符串。
[要求]
如果strArr长度为N，时间复杂度请达到O(NlogK)

输出K行，每行有一个字符串和一个整数（字符串表示）。
你需要按照出现出现次数由大到小输出，若出现次数相同时字符串字典序较小的优先输出

示例1
输入
复制
["1","2","3","4"],2
返回值
复制
[["1","1"],["2","1"]]
示例2
输入
复制
["1","1","2","3"],2
返回值
复制
[["1","2"],["2","1"]]
*/

//采用堆方法(xi)
func topKstrings(strings []string, k int) [][]string {
	// write code here
	if len(strings) == 0 {
		return [][]string{}
	}

	times := make(map[string]int64)
	for i := range strings {
		times[strings[i]] = times[strings[i]] + 1
	}

	h := &heapString{}
	heap.Init(h)

	//遍历times
	for key, val := range times {
		if h.Len() < k {
			heap.Push(h, tString{Val: key, time: val})
		} else if (*h)[0].time < val || ((*h)[0].time == val && (*h)[0].Val > key) {
			(*h)[0].time = val
			(*h)[0].Val = key
			heap.Fix(h, 0)
		}
	}

	//fmt.Println(*h)
	////遍历堆中的k个元素
	//sort.Slice(*h, func(i, j int) bool {
	//	if (*h)[i].time > (*h)[j].time {
	//		return true
	//	} else if (*h)[i].time == (*h)[j].time {
	//		return (*h)[i].Val > (*h)[j].Val
	//	}
	//
	//	return false
	//})
	//
	//fmt.Println(*h)

	//[{2 1} {1 2}]
	ret := make([][]string, k)
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(h).(tString)
		tmp := make([]string, 2)
		tmp[0] = item.Val
		tmp[1] = strconv.FormatInt(item.time, 10)
		ret[i] = tmp
	}

	return ret
}

type tString struct {
	Val  string
	time int64
}

type heapString []tString

func (h heapString) Len() int { return len(h) }

// 堆顶是出现次数最多的
func (h heapString) Less(i, j int) bool {
	if h[i].time != h[j].time {
		return h[i].time < h[j].time
	}
	return h[i].Val > h[j].Val
}

func (h heapString) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapString) Push(x interface{}) {
	*h = append(*h, x.(tString))
}

func (h *heapString) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	//arr := []int{9, 1, 3, 2, 4, 6}
	//fmt.Println(findKthMin(arr, 3))
	//fmt.Println(GetLeastNumbers_Solution2([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4))
	fmt.Println(topKstrings([]string{"1", "1", "2", "3"}, 2))
}
