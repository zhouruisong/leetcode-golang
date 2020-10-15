package main

import (
	"main/model"
)

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。
*/


func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	//新增节点，防止删除的节点为头节点，
	newHead := &model.ListNode{
		Next: head,
	}

	first := head
	pre := newHead
	preNode := pre
	offset := 0
	for first != nil {
		first = first.Next
		if offset < n {
			offset++
			continue
		}
		preNode = preNode.Next
	}

	delNode := preNode.Next
	preNode.Next = delNode.Next
	delNode.Next = nil
	return pre.Next
}

func main() {
	v1 := []int{1, 2}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	head = removeNthFromEnd(head, 1)
	model.PrintList(head)
}
