package main

import "main/model"

/*
插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。


示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

//o(N)-o(N*N)
func insertionSortList(head *model.ListNode) *model.ListNode {
	if head == nil {
		return head
	}

	newHead := &model.ListNode{}
	preloc := newHead
	loc := newHead.Next

	start := head
	for start != nil {
		next := start.Next
		start.Next = nil

		//查询插入位置
		for loc != nil {
			if loc.Val >= start.Val {
				break
			}
			preloc = loc
			loc = loc.Next
		}
		preloc.Next = start
		start.Next = loc

		preloc = newHead
		loc = newHead.Next
		start = next
	}

	return newHead.Next
}

func main() {
	v1 := []int{-1, 5, 3, 4, 0}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	head = insertionSortList(head)
	model.PrintList(head)
}
