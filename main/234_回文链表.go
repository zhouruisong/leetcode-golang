package main

import "main/model"

/*
* 请判断一个链表是否为回文链表。
*
* 示例 1:
*
* 输入: 1->2
* 输出: false
*
* 示例 2:
*
* 输入: 1->2->2->1
* 输出: true
*
*
* 进阶：
* 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *model.ListNode) bool {
	if head == nil {
		return true
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var pre, cur *model.ListNode = nil, slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	mid := pre
	for mid != nil {
		if head.Val != mid.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}

	return true
}
