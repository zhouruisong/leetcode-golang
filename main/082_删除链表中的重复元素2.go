package main

import "main/model"

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

func deleteDuplicates(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//哨兵res
	res := new(model.ListNode)
	p := res
	//双指针
	curr := head.Next
	pre := head

	for curr != nil {
		if pre.Val == curr.Val {
			for curr != nil && curr.Val == pre.Val {
				//找到第一个与pre不重复的节点
				curr = curr.Next
			}
			if curr == nil {
				//从pre开始到结束都是重复的节点
				p.Next = nil
				return res.Next
			}
			//更新pre和curr
			pre = curr
			curr = curr.Next
			//进入下一次判断
			continue
		} else {
			p.Next = pre
			curr = curr.Next
			pre = pre.Next
			p = p.Next
		}
	}

	if p.Val != pre.Val {
		p.Next = pre
		p = p.Next
		p.Next = nil
	}

	return res.Next
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 5}
	tmp1 := model.UnmarshalListBySlice(x)
	model.PrintList(tmp1)
	tmp2 := deleteDuplicates(tmp1)
	model.PrintList(tmp2)
}
