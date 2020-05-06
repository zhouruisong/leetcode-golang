package main

import "main/model"

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/

func partition(head *model.ListNode, x int) *model.ListNode {
	if head == nil {
		return nil
	}

	beforePre := new(model.ListNode)
	afterPre := new(model.ListNode)
	beforeNode := beforePre
	afterNode := afterPre

	for head != nil {
		if head.Val < x {
			beforeNode.Next = head
			beforeNode = beforeNode.Next
		} else {
			afterNode.Next = head
			afterNode = afterNode.Next
		}
		head = head.Next
	}

	afterNode.Next = nil
	//model.PrintList(afterPre.Next)
	//model.PrintList(beforePre.Next)

	beforeNode.Next = afterPre.Next
	return beforePre.Next
}

func main() {
	x := []int{1, 4, 3, 2, 5, 2}
	tmp1 := model.UnmarshalListBySlice(x)
	model.PrintList(tmp1)
	tmp2 := partition(tmp1, 3)
	model.PrintList(tmp2)
}
