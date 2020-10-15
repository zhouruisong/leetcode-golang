package main

import "container/heap"

/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

 

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
 

提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
 */

func topKFrequent(nums []int, k int) []int {
	/*方法2：使用大根堆, 在map统计完词频后, 放到切片中, 然后用heap.init方法(该方法原理是从底向上建堆,
	时间复杂度是O(N), N是切片的长度, 也就是统计词频表map的长度), 然后pop, 一共弹出topk次, 所以该种解法的时间复杂度是O(topk * logn)
	*/
	if k == 0 || len(nums) == 0 {
		return make([]int, 0)
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	nodes := make(NodeHeap, 0, len(m))
	for k, v := range m {
		nodes = append(nodes, &Node{
			val:   k,
			times: v,
		})
	}
	h := &nodes
	heap.Init(h)
	topK := min(k, len(m))
	res := make([]int, 0, topK)
	for i := 0; i < topK; i++ {
		res = append(res, heap.Pop(h).(*Node).val)
	}
	return res
}

type Node struct {
	val   int
	times int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int { return len(h) }

// 大根堆
func (h NodeHeap) Less(i, j int) bool { return h[i].times > h[j].times }

func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
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
