package main

import (
	"fmt"
	"leetcode-golang/main/model"
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
	nodeList := []*model.ListNode{}

	for first != nil {
		nodeList = append(nodeList, first)
		first = first.Next
		count++
		if count%2 == 0 {
			for len(nodeList) > 0 {
				b := nodeList[len(nodeList)-1]
				nodeList = nodeList[:len(nodeList)-1]
				//fmt.Println(b.Val)
				tmp.Next = b
				tmp = tmp.Next
			}
		}
	}

	for len(nodeList) > 0 {
		//fmt.Println(b.Val)
		b := nodeList[len(nodeList)-1]
		nodeList = nodeList[:len(nodeList)-1]
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

func reverseKGroup(head *model.ListNode, k int) *model.ListNode {
	/*
		首先通过k次for循环用end指针指向翻转元素的末尾
		此时判断一下如果翻转元素不到k个，即end==null，说明已经到达末尾，直接返回即可
		接下来需要定义两个指针pre和pLast分别记录翻转元素的前驱和后继，以便将翻转元素前后两部分连接起来
		之后再重置pre和end 指针，进入下一次循环
		遍历完之后返回dummy带头结点接下来的元素即可。
	*/
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	newHead := &model.ListNode{
		Next: head,
	}

	pre := newHead
	end := pre
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		next := end.Next
		end.Next = nil
		//反转链表，范围是[pre,end]
		cur := pre.Next
		pre.Next = reverseLink(cur)
		//从新连接起来
		cur.Next = next
		pre = cur
		end = pre
	}

	return newHead.Next
}

func reverseLink(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var tmp *model.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = tmp
		tmp = cur
		cur = next
	}

	return tmp
}

func main() {
	v1 := []int{1, 2, 3, 4, 5, 6}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	fmt.Println("=======================")
	//nh := swapPairs3(head)
	//借助k个一组旋转链表
	nh := reverseKGroup(head, 2)
	model.PrintList(nh)
}
