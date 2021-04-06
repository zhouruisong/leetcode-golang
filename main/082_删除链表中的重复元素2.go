package main

import "main/model"

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

func deleteDuplicates2(head *model.ListNode) *model.ListNode {
	dummy := &model.ListNode{Next: head}
	tmp := dummy
	slow, fast := head, head
	for fast != nil {
		for fast.Next != nil && fast.Next.Val == fast.Val {
			fast = fast.Next
		}
		if slow == fast {
			tmp.Next = slow
			tmp = tmp.Next
		}
		fast = fast.Next
		slow = fast
	}

	tmp.Next = nil
	return dummy.Next
}

func main() {
	x := []int{1, 2, 3, 3, 4, 4, 5}
	tmp1 := model.UnmarshalListBySlice(x)
	model.PrintList(tmp1)
	tmp3 := deleteDuplicates2(tmp1)
	model.PrintList(tmp3)
}
