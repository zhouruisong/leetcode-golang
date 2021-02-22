package main

import "main/model"

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

func reverseBetween(head *model.ListNode, m int, n int) *model.ListNode {
	if head == nil {
		return nil
	}

	newHead := &model.ListNode{
		Next: head,
	}

	pre := newHead
	start := 0
	for pre != nil {
		if start < m-1 {
			pre = pre.Next
			start++
			continue
		}
		//找到m的前一个位置
		p := pre.Next
		pre.Next = nil
		var tmp *model.ListNode
		//这个范围链表反转
		//从p开始反转链表,start == n 时结束
		for start < n && p != nil {
			next := p.Next
			p.Next = tmp
			tmp = p
			p = next
			start++
		}

		//追加反转后的数据
		pre.Next = tmp
		for pre.Next != nil {
			pre = pre.Next
		}
		pre.Next = p
		break
	}

	return newHead.Next
}

func add_node_head(head, newNode *model.ListNode) {
	if head == nil || newNode == nil {
		return
	}

	newNode.Next = head.Next
	head.Next = newNode
}

func main() {
	v6 := []int{1, 2, 3, 4, 5, 6, 7}
	tmp1 := model.UnmarshalListBySlice(v6)
	model.PrintList(tmp1)
	tmp2 := reverseBetween(tmp1, 2, 4)
	model.PrintList(tmp2)
}
