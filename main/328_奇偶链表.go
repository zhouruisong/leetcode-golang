package main

import "main/model"

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/

func oddEvenList(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}

	head1, head2 := head, head.Next
	if head2 == nil {
		return head
	}

	cur := head2.Next
	if cur == nil {
		return head
	}

	tmp1 := head1
	tmp2 := head2

	n := 3
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		//奇数
		if n%2 != 0 {
			head1.Next = cur
			head1 = head1.Next
		} else {
			head2.Next = cur
			head2 = head2.Next
		}
		cur = next
		n++
	}

	head1.Next = tmp2
	head2.Next = nil

	return tmp1
}

func main() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	head = oddEvenList(head)
	model.PrintList(head)
}
