package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func main() {
	//arr := []int{9, 1, 3, 2, 4, 6}
	//fmt.Println(findKthMin(arr, 3))
	fmt.Println(GetLeastNumbers_Solution2([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4))
}
