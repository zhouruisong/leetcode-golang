package main

import "leetcode-golang/main/model"

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */

/*
这里先假设链表A头结点与结点8的长度 与 链表B头结点与结点8的长度相等，那么就可以用双指针。

初始化：指针ta指向链表A头结点，指针tb指向链表B头结点
如果ta == tb， 说明找到了第一个公共的头结点，直接返回即可。
否则，ta != tb，则++ta，++tb
所以现在的问题就变成，如何让本来长度不相等的变为相等的？
假设链表A长度为a， 链表B的长度为b，此时a != b
但是，a+b == b+a
因此，可以让a+b作为链表A的新长度，b+a作为链表B的新长度。
*/
func FindFirstCommonNode(pHead1 *model.ListNode, pHead2 *model.ListNode) *model.ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	p1 := pHead1
	p2 := pHead2
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 != p2 {
			if p1 == nil {
				p1 = pHead2
			}
			if p2 == nil {
				p2 = pHead1
			}
		}
	}

	return p1
}
