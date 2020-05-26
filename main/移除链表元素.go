package main

import "main/model"

/*
 * 删除链表中等于给定值 val 的所有节点。
 *
 * 示例:
 *
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
 */

func removeElements(head *model.ListNode, val int) *model.ListNode {
	if head == nil {
		return nil
	}

	newhead := &model.ListNode{Next: head}
	cur := newhead.Next
	pre := newhead
	for cur != nil {
		if cur.Val == val {
			cur = cur.Next
		} else {
			pre.Next = cur
			pre = pre.Next
			cur = cur.Next
		}
	}

	pre.Next = cur
	return newhead.Next
}

func main() {
	v6 := []int{6, 6, 6}
	tmp := model.UnmarshalListBySlice(v6)
	tmp2 := removeElements(tmp, 6)
	model.PrintList(tmp2)
}
