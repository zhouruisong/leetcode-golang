package main

import "practies/model"

/*
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *model.ListNode) *model.ListNode {
	p1 := head
	p2 := head
	tmp := p2
	for p1 != nil {
		p := p1.Next
		for p != nil && p.Val == p2.Val {
			p = p.Next
		}

		p2.Next = p
		p2 = p2.Next
		p1 = p2
	}

	return tmp
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 5, 5}
	//x := []int{1, 1, 1}
	//x := []int{1, 1, 1}

	tmp1 := model.UnmarshalListBySlice(x)
	model.PrintList(tmp1)
	tmp2 := deleteDuplicates(tmp1)
	model.PrintList(tmp2)
}
