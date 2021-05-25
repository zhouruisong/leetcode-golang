package main

import "leetcode-golang/main/model"

/*
* 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 *
 *
 * 示例：
 *
 * 给你这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 *
 *
 * 说明：
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/*
首先通过k次for循环用end指针指向翻转元素的末尾

此时判断一下如果翻转元素不到k个，即end==null，说明已经到达末尾，直接返回即可

接下来需要定义两个指针pre和pLast分别记录翻转元素的前驱和后继，以便将翻转元素前后两部分连接起来

之后再重置pre和end 指针，进入下一次循环

遍历完之后返回dummy带头结点接下来的元素即可。
*/

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
	v6 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tmp1 := model.UnmarshalListBySlice(v6)
	model.PrintList(tmp1)
	tmp2 := reverseKGroup(tmp1, 2)
	model.PrintList(tmp2)
}
