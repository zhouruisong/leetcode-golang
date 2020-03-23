package main

import (
	"main/model"
)

func mergeTwoLists(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	var head *model.ListNode
	start := new(model.ListNode)
	head = start

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			start.Next = l1
			l1 = l1.Next
			start = start.Next
		} else {
			start.Next = l2
			l2 = l2.Next
			start = start.Next
		}
	}

	if l1 != nil && l2 == nil {
		start.Next = l1
	}

	if l1 == nil && l2 != nil {
		start.Next = l2
	}

	return head.Next
}

func mergeTwoLists2(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	p1 := list1
	p2 := list2
	head := p2
	if p1.Val < p2.Val {
		head = p1
		p1 = p1.Next
	} else {
		head = p2
		p2 = p2.Next
	}

	tmp := head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}

	if p1 != nil {
		head.Next = p1
	}

	if p2 != nil {
		head.Next = p2
	}

	return tmp
}

func main() {
	v1 := []int{1, 2, 5}
	v2 := []int{2, 4, 5, 6, 8}
	tmp1 := model.UnmarshalListBySlice(v1)
	tmp2 := model.UnmarshalListBySlice(v2)
	model.PrintList(tmp1)
	model.PrintList(tmp2)

	tmp3 := mergeTwoLists2(tmp1, tmp2)

	model.PrintList(tmp3)
}
