package main

import "main/model"

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

 

进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

 

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	//反转
	l1 = reserveLink(l1)
	l2 = reserveLink(l2)

	result := new(model.ListNode)
	head := result
	sum := 0 //进位
	tmp := 0
	v := 0 //本位数
	for l1 != nil && l2 != nil {
		tmp = sum + l1.Val + l2.Val //对位累加后再加上后面的进位
		v = tmp % 10                //本位数
		//fmt.Println(tmp)
		sum = tmp / 10 //进位
		node := &model.ListNode{
			Val: v,
		}
		result.Next = node
		result = result.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	//不相等链表
	if l1 == nil && l2 != nil {
		for l2 != nil {
			tmp = sum + l2.Val
			v = tmp % 10
			sum = tmp / 10
			node := &model.ListNode{
				Val: v,
			}
			result.Next = node
			result = result.Next
			l2 = l2.Next
		}
	}

	if l1 != nil && l2 == nil {
		for l1 != nil {
			tmp = sum + l1.Val
			v = tmp % 10
			sum = tmp / 10
			node := &model.ListNode{
				Val: v,
			}
			result.Next = node
			result = result.Next
			l1 = l1.Next
		}
	}

	//最高位是1需要进位
	if sum > 0 {
		node := &model.ListNode{
			Val: sum,
		}
		result.Next = node
	}

	head.Next = reserveLink(head.Next)
	return head.Next
}

func reserveLink(head *model.ListNode) *model.ListNode {
	cur := head
	var p *model.ListNode
	var next *model.ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = p
		p = cur
		cur = next
	}
	return p
}
