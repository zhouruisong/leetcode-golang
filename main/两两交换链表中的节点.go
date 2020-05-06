package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

func swapPairs(head *model.ListNode) *model.ListNode {
	//借助栈
	if head == nil {
		return head
	}

	newHead1 := &model.ListNode{}
	tmp := newHead1

	first := head
	count := 0
	nodeList := list.New()

	for first != nil {
		nodeList.PushBack(first)
		first = first.Next
		count++
		if count%2 == 0 {
			for nodeList.Len() > 0 {
				b := nodeList.Remove(nodeList.Back()).(*model.ListNode)
				//fmt.Println(b.Val)
				tmp.Next = b
				tmp = tmp.Next
			}
		}
	}

	for nodeList.Len() > 0 {
		b := nodeList.Remove(nodeList.Back()).(*model.ListNode)
		//fmt.Println(b.Val)
		tmp.Next = b
		tmp = tmp.Next
	}
	tmp.Next = nil
	return newHead1.Next
}

//原地交换
func swapPairs2(head *model.ListNode) *model.ListNode {
	if head == nil {
		return head
	}

	newHead := &model.ListNode{
		Next: head,
	}
	pre := newHead

	//a1->a2->a3->a4
	for pre.Next != nil && pre.Next.Next != nil {
		a1 := pre
		a2 := pre.Next
		a3 := pre.Next.Next
		a4 := pre.Next.Next.Next

		a1.Next = a3
		a2.Next = a4
		a3.Next = a2
		pre = pre.Next.Next
	}

	return newHead.Next
}

//递归
func swapPairs3(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil { // 终止条件
		return head
	}
	firstNode := head
	secondNode := head.Next
	firstNode.Next = swapPairs(secondNode.Next) // 调用递归
	secondNode.Next = firstNode
	return secondNode
}

func main() {
	v1 := []int{1, 2, 3, 4, 5, 6}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	fmt.Println("=======")
	nh := swapPairs3(head)
	model.PrintList(nh)
}
