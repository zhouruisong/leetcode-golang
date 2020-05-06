package main

import "main/model"

/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/

func reorderList(head *model.ListNode) {
	if head == nil {
		return
	}

	//先找链表的中间节点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left := head
	right := slow.Next
	//左右分开
	slow.Next = nil

	//将right反转
	right = reverseLink(right)

	for left != nil && right != nil {
		tmp1 := left.Next
		tmp2 := right.Next
		left.Next = right
		right.Next = tmp1
		left = tmp1
		right = tmp2
	}

	return
}

func reverseLink(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}

	var pre *model.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	tmp1 := model.UnmarshalListBySlice(x)
	model.PrintList(tmp1)
	reorderList(tmp1)
	model.PrintList(tmp1)
}
