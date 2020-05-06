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
	v1 := []int{7, 2, 4, 3}
	v2 := []int{5, 6, 4}
	//[7,8,0,7]
	tmp := model.UnmarshalListBySlice(v1)
	tmp2 := model.UnmarshalListBySlice(v2)
	model.PrintList(tmp)
	model.PrintList(tmp2)

	ret := addTwoNumbers(tmp, tmp2)
	model.PrintList(ret)
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

func addTwoNumbers(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
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
