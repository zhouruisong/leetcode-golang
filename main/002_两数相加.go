package main

import (
	"main/model"
)

/*
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func main() {
	v1 := []int{2, 4, 9}
	v2 := []int{5, 6, 4, 9}
	//[7,8,0,7]

	tmp3 := model.UnmarshalListBySlice(v1)
	tmp4 := model.UnmarshalListBySlice(v2)
	model.PrintList(tmp3)
	model.PrintList(tmp4)
	ret2 := addTwoNumbers2(tmp3, tmp4)
	model.PrintList(ret2)
}

func addTwoNumbers2(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	result := &model.ListNode{}
	head := result
	carry := 0
	for l1 != nil || l2 != nil {
		node := &model.ListNode{}
		if l1 == nil {
			node.Val = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		} else if l2 == nil {
			node.Val = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
		} else {
			node.Val = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		}

		result.Next = node
		node.Next = nil
		result = result.Next
	}

	if carry > 0 {
		node := &model.ListNode{
			Val: carry,
		}
		result.Next = node
		node.Next = nil
		result = result.Next
	}

	return head.Next
}
