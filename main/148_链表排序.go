package main

import "main/model"

/*
 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 *
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 *
*/

func sortList(head *model.ListNode) *model.ListNode {
	// O(n log n)采用归并排序
	if head == nil || head.Next == nil {
		return head
	}

	//找到中间节点
	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left := head
	right := slow.Next
	slow.Next = nil

	left = sortList(left)
	right = sortList(right)
	return merge(left, right)
}

func merge(head1, head2 *model.ListNode) *model.ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	p1 := head1
	p2 := head2
	head := head2
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
			head = head.Next
		} else {
			head.Next = p2
			p2 = p2.Next
			head = head.Next
		}
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
	v1 := []int{5, 2, 1, 3, 4, 9, 8, 7, 6}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	head = sortList(head)
	model.PrintList(head)
}
